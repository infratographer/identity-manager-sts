// Package httpsrv provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.5-0.20230118012357-f4cf8f9a5703 DO NOT EDIT.
package httpsrv

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	. "go.infratographer.com/identity-manager-sts/pkg/api/v1"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Deletes an issuer with the given ID.
	// (DELETE /api/v1/issuers/{id})
	DeleteIssuer(c *gin.Context, id openapi_types.UUID)
	// Gets an issuer by ID.
	// (GET /api/v1/issuers/{id})
	GetIssuerByID(c *gin.Context, id openapi_types.UUID)
	// Updates an issuer.
	// (PATCH /api/v1/issuers/{id})
	UpdateIssuer(c *gin.Context, id openapi_types.UUID)
	// Creates an issuer.
	// (POST /api/v1/tenants/{tenantID}/issuers)
	CreateIssuer(c *gin.Context, tenantID openapi_types.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// DeleteIssuer operation middleware
func (siw *ServerInterfaceWrapper) DeleteIssuer(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteIssuer(c, id)
}

// GetIssuerByID operation middleware
func (siw *ServerInterfaceWrapper) GetIssuerByID(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetIssuerByID(c, id)
}

// UpdateIssuer operation middleware
func (siw *ServerInterfaceWrapper) UpdateIssuer(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateIssuer(c, id)
}

// CreateIssuer operation middleware
func (siw *ServerInterfaceWrapper) CreateIssuer(c *gin.Context) {

	var err error

	// ------------- Path parameter "tenantID" -------------
	var tenantID openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "tenantID", c.Param("tenantID"), &tenantID)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter tenantID: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateIssuer(c, tenantID)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.DELETE(options.BaseURL+"/api/v1/issuers/:id", wrapper.DeleteIssuer)
	router.GET(options.BaseURL+"/api/v1/issuers/:id", wrapper.GetIssuerByID)
	router.PATCH(options.BaseURL+"/api/v1/issuers/:id", wrapper.UpdateIssuer)
	router.POST(options.BaseURL+"/api/v1/tenants/:tenantID/issuers", wrapper.CreateIssuer)
}

type DeleteIssuerRequestObject struct {
	Id openapi_types.UUID `json:"id"`
}

type DeleteIssuerResponseObject interface {
	VisitDeleteIssuerResponse(w http.ResponseWriter) error
}

type DeleteIssuer200JSONResponse DeleteResponse

func (response DeleteIssuer200JSONResponse) VisitDeleteIssuerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteIssuer404JSONResponse ErrorResponse

func (response DeleteIssuer404JSONResponse) VisitDeleteIssuerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetIssuerByIDRequestObject struct {
	Id openapi_types.UUID `json:"id"`
}

type GetIssuerByIDResponseObject interface {
	VisitGetIssuerByIDResponse(w http.ResponseWriter) error
}

type GetIssuerByID200JSONResponse Issuer

func (response GetIssuerByID200JSONResponse) VisitGetIssuerByIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetIssuerByID404JSONResponse ErrorResponse

func (response GetIssuerByID404JSONResponse) VisitGetIssuerByIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type UpdateIssuerRequestObject struct {
	Id   openapi_types.UUID `json:"id"`
	Body *UpdateIssuerJSONRequestBody
}

type UpdateIssuerResponseObject interface {
	VisitUpdateIssuerResponse(w http.ResponseWriter) error
}

type UpdateIssuer200JSONResponse Issuer

func (response UpdateIssuer200JSONResponse) VisitUpdateIssuerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateIssuer400JSONResponse ErrorResponse

func (response UpdateIssuer400JSONResponse) VisitUpdateIssuerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type UpdateIssuer404JSONResponse ErrorResponse

func (response UpdateIssuer404JSONResponse) VisitUpdateIssuerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type CreateIssuerRequestObject struct {
	TenantID openapi_types.UUID `json:"tenantID"`
	Body     *CreateIssuerJSONRequestBody
}

type CreateIssuerResponseObject interface {
	VisitCreateIssuerResponse(w http.ResponseWriter) error
}

type CreateIssuer200JSONResponse Issuer

func (response CreateIssuer200JSONResponse) VisitCreateIssuerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateIssuer400JSONResponse ErrorResponse

func (response CreateIssuer400JSONResponse) VisitCreateIssuerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Deletes an issuer with the given ID.
	// (DELETE /api/v1/issuers/{id})
	DeleteIssuer(ctx context.Context, request DeleteIssuerRequestObject) (DeleteIssuerResponseObject, error)
	// Gets an issuer by ID.
	// (GET /api/v1/issuers/{id})
	GetIssuerByID(ctx context.Context, request GetIssuerByIDRequestObject) (GetIssuerByIDResponseObject, error)
	// Updates an issuer.
	// (PATCH /api/v1/issuers/{id})
	UpdateIssuer(ctx context.Context, request UpdateIssuerRequestObject) (UpdateIssuerResponseObject, error)
	// Creates an issuer.
	// (POST /api/v1/tenants/{tenantID}/issuers)
	CreateIssuer(ctx context.Context, request CreateIssuerRequestObject) (CreateIssuerResponseObject, error)
}

type StrictHandlerFunc func(ctx *gin.Context, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// DeleteIssuer operation middleware
func (sh *strictHandler) DeleteIssuer(ctx *gin.Context, id openapi_types.UUID) {
	var request DeleteIssuerRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteIssuer(ctx, request.(DeleteIssuerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteIssuer")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(DeleteIssuerResponseObject); ok {
		if err := validResponse.VisitDeleteIssuerResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// GetIssuerByID operation middleware
func (sh *strictHandler) GetIssuerByID(ctx *gin.Context, id openapi_types.UUID) {
	var request GetIssuerByIDRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetIssuerByID(ctx, request.(GetIssuerByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetIssuerByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(GetIssuerByIDResponseObject); ok {
		if err := validResponse.VisitGetIssuerByIDResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// UpdateIssuer operation middleware
func (sh *strictHandler) UpdateIssuer(ctx *gin.Context, id openapi_types.UUID) {
	var request UpdateIssuerRequestObject

	request.Id = id

	var body UpdateIssuerJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateIssuer(ctx, request.(UpdateIssuerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateIssuer")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(UpdateIssuerResponseObject); ok {
		if err := validResponse.VisitUpdateIssuerResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// CreateIssuer operation middleware
func (sh *strictHandler) CreateIssuer(ctx *gin.Context, tenantID openapi_types.UUID) {
	var request CreateIssuerRequestObject

	request.TenantID = tenantID

	var body CreateIssuerJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateIssuer(ctx, request.(CreateIssuerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateIssuer")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(CreateIssuerResponseObject); ok {
		if err := validResponse.VisitCreateIssuerResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}
