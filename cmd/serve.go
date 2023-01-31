package cmd

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ory/fosite/compose"
	fositestorage "github.com/ory/fosite/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.infratographer.com/x/ginx"
	"go.infratographer.com/x/otelx"
	"go.uber.org/zap/zapcore"

	"go.infratographer.com/identity-manager-sts/internal/api/httpsrv"
	"go.infratographer.com/identity-manager-sts/internal/config"
	"go.infratographer.com/identity-manager-sts/internal/fositex"
	"go.infratographer.com/identity-manager-sts/internal/jwks"
	"go.infratographer.com/identity-manager-sts/internal/rfc8693"
	"go.infratographer.com/identity-manager-sts/internal/routes"
	"go.infratographer.com/identity-manager-sts/internal/storage"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts identity-manager-sts",
	Run: func(cmd *cobra.Command, args []string) {
		serve(cmd.Context())
	},
}

var (
	defaultListen = ":8080"
)

func init() {
	rootCmd.AddCommand(serveCmd)

	ginx.MustViperFlags(viper.GetViper(), serveCmd.Flags(), defaultListen)
	otelx.MustViperFlags(viper.GetViper(), serveCmd.Flags())
}

func serve(ctx context.Context) {
	err := otelx.InitTracer(config.Config.OTel, appName, logger)
	if err != nil {
		logger.Fatalf("error initializing tracing: %s", err)
	}

	storageEngine, err := storage.NewEngine(config.Config.Storage)
	if err != nil {
		logger.Fatalf("error initializing storage: %s", err)
	}

	defer storageEngine.Shutdown()

	mappingStrategy := rfc8693.NewClaimMappingStrategy(storageEngine)

	jwksStrategy := jwks.NewIssuerJWKSURIStrategy(storageEngine)

	oauth2Config, err := fositex.NewOAuth2Config(config.Config.OAuth)
	if err != nil {
		logger.Fatalf("error loading config: %s", err)
	}

	oauth2Config.IssuerJWKSURIStrategy = jwksStrategy
	oauth2Config.ClaimMappingStrategy = mappingStrategy

	keyGetter := func(ctx context.Context) (any, error) {
		return oauth2Config.GetSigningKey(ctx), nil
	}

	hmacStrategy := compose.NewOAuth2HMACStrategy(oauth2Config)
	jwtStrategy := compose.NewOAuth2JWTStrategy(keyGetter, hmacStrategy, oauth2Config)
	store := fositestorage.NewExampleStore()
	tokenExchangeHandler := rfc8693.NewTokenExchangeHandler(oauth2Config, jwtStrategy, store)
	oauth2Config.TokenEndpointHandlers.Append(tokenExchangeHandler)
	provider := fositex.NewOAuth2Provider(oauth2Config, store)

	apiHandler, err := httpsrv.NewAPIHandler(storageEngine)
	if err != nil {
		logger.Fatal("error initializing API server: %s", err)
	}

	router := routes.NewRouter(logger, oauth2Config, provider)

	emptyLogFn := func(c *gin.Context) []zapcore.Field {
		return []zapcore.Field{}
	}

	// ginx doesn't allow configuration of ContextWithFallback but we need it here.
	engine := ginx.DefaultEngine(logger.Desugar(), emptyLogFn)
	engine.ContextWithFallback = true

	router.Routes(engine.Group("/"))
	apiHandler.Routes(engine.Group("/"))

	srv := &http.Server{
		Addr:    config.Config.Server.Listen,
		Handler: engine,
	}

	logger.Fatal(srv.ListenAndServe())
}
