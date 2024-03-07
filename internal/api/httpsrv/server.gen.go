// Package httpsrv provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package httpsrv

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	. "go.infratographer.com/identity-api/pkg/api/v1"
	"go.infratographer.com/x/gidx"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Deletes an OAuth Client
	// (DELETE /api/v1/clients/{clientID})
	DeleteOAuthClient(ctx echo.Context, clientID gidx.PrefixedID) error
	// Gets information about an OAuth 2.0 Client.
	// (GET /api/v1/clients/{clientID})
	GetOAuthClient(ctx echo.Context, clientID gidx.PrefixedID) error
	// Deletes an issuer with the given ID.
	// (DELETE /api/v1/issuers/{id})
	DeleteIssuer(ctx echo.Context, id gidx.PrefixedID) error
	// Gets an issuer by ID.
	// (GET /api/v1/issuers/{id})
	GetIssuerByID(ctx echo.Context, id gidx.PrefixedID) error
	// Updates an issuer.
	// (PATCH /api/v1/issuers/{id})
	UpdateIssuer(ctx echo.Context, id gidx.PrefixedID) error
	// Creates an OAuth client.
	// (POST /api/v1/owners/{ownerID}/clients)
	CreateOAuthClient(ctx echo.Context, ownerID gidx.PrefixedID) error
	// Creates an issuer.
	// (POST /api/v1/owners/{ownerID}/issuers)
	CreateIssuer(ctx echo.Context, ownerID gidx.PrefixedID) error
	// Gets information about a User.
	// (GET /api/v1/users/{userID})
	GetUserByID(ctx echo.Context, userID gidx.PrefixedID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// DeleteOAuthClient converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteOAuthClient(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clientID" -------------
	var clientID gidx.PrefixedID

	err = runtime.BindStyledParameterWithOptions("simple", "clientID", ctx.Param("clientID"), &clientID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clientID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteOAuthClient(ctx, clientID)
	return err
}

// GetOAuthClient converts echo context to params.
func (w *ServerInterfaceWrapper) GetOAuthClient(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clientID" -------------
	var clientID gidx.PrefixedID

	err = runtime.BindStyledParameterWithOptions("simple", "clientID", ctx.Param("clientID"), &clientID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clientID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetOAuthClient(ctx, clientID)
	return err
}

// DeleteIssuer converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteIssuer(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id gidx.PrefixedID

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteIssuer(ctx, id)
	return err
}

// GetIssuerByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetIssuerByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id gidx.PrefixedID

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetIssuerByID(ctx, id)
	return err
}

// UpdateIssuer converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateIssuer(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id gidx.PrefixedID

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateIssuer(ctx, id)
	return err
}

// CreateOAuthClient converts echo context to params.
func (w *ServerInterfaceWrapper) CreateOAuthClient(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "ownerID" -------------
	var ownerID gidx.PrefixedID

	err = runtime.BindStyledParameterWithOptions("simple", "ownerID", ctx.Param("ownerID"), &ownerID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ownerID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateOAuthClient(ctx, ownerID)
	return err
}

// CreateIssuer converts echo context to params.
func (w *ServerInterfaceWrapper) CreateIssuer(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "ownerID" -------------
	var ownerID gidx.PrefixedID

	err = runtime.BindStyledParameterWithOptions("simple", "ownerID", ctx.Param("ownerID"), &ownerID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ownerID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateIssuer(ctx, ownerID)
	return err
}

// GetUserByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userID" -------------
	var userID gidx.PrefixedID

	err = runtime.BindStyledParameterWithOptions("simple", "userID", ctx.Param("userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserByID(ctx, userID)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE(baseURL+"/api/v1/clients/:clientID", wrapper.DeleteOAuthClient)
	router.GET(baseURL+"/api/v1/clients/:clientID", wrapper.GetOAuthClient)
	router.DELETE(baseURL+"/api/v1/issuers/:id", wrapper.DeleteIssuer)
	router.GET(baseURL+"/api/v1/issuers/:id", wrapper.GetIssuerByID)
	router.PATCH(baseURL+"/api/v1/issuers/:id", wrapper.UpdateIssuer)
	router.POST(baseURL+"/api/v1/owners/:ownerID/clients", wrapper.CreateOAuthClient)
	router.POST(baseURL+"/api/v1/owners/:ownerID/issuers", wrapper.CreateIssuer)
	router.GET(baseURL+"/api/v1/users/:userID", wrapper.GetUserByID)

}

type DeleteOAuthClientRequestObject struct {
	ClientID gidx.PrefixedID `json:"clientID"`
}

type DeleteOAuthClientResponseObject interface {
	VisitDeleteOAuthClientResponse(w http.ResponseWriter) error
}

type DeleteOAuthClient200JSONResponse DeleteResponse

func (response DeleteOAuthClient200JSONResponse) VisitDeleteOAuthClientResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetOAuthClientRequestObject struct {
	ClientID gidx.PrefixedID `json:"clientID"`
}

type GetOAuthClientResponseObject interface {
	VisitGetOAuthClientResponse(w http.ResponseWriter) error
}

type GetOAuthClient200JSONResponse OAuthClient

func (response GetOAuthClient200JSONResponse) VisitGetOAuthClientResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteIssuerRequestObject struct {
	Id gidx.PrefixedID `json:"id"`
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

type GetIssuerByIDRequestObject struct {
	Id gidx.PrefixedID `json:"id"`
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

type UpdateIssuerRequestObject struct {
	Id   gidx.PrefixedID `json:"id"`
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

type CreateOAuthClientRequestObject struct {
	OwnerID gidx.PrefixedID `json:"ownerID"`
	Body    *CreateOAuthClientJSONRequestBody
}

type CreateOAuthClientResponseObject interface {
	VisitCreateOAuthClientResponse(w http.ResponseWriter) error
}

type CreateOAuthClient200JSONResponse OAuthClient

func (response CreateOAuthClient200JSONResponse) VisitCreateOAuthClientResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateIssuerRequestObject struct {
	OwnerID gidx.PrefixedID `json:"ownerID"`
	Body    *CreateIssuerJSONRequestBody
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

type GetUserByIDRequestObject struct {
	UserID gidx.PrefixedID `json:"userID"`
}

type GetUserByIDResponseObject interface {
	VisitGetUserByIDResponse(w http.ResponseWriter) error
}

type GetUserByID200JSONResponse User

func (response GetUserByID200JSONResponse) VisitGetUserByIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Deletes an OAuth Client
	// (DELETE /api/v1/clients/{clientID})
	DeleteOAuthClient(ctx context.Context, request DeleteOAuthClientRequestObject) (DeleteOAuthClientResponseObject, error)
	// Gets information about an OAuth 2.0 Client.
	// (GET /api/v1/clients/{clientID})
	GetOAuthClient(ctx context.Context, request GetOAuthClientRequestObject) (GetOAuthClientResponseObject, error)
	// Deletes an issuer with the given ID.
	// (DELETE /api/v1/issuers/{id})
	DeleteIssuer(ctx context.Context, request DeleteIssuerRequestObject) (DeleteIssuerResponseObject, error)
	// Gets an issuer by ID.
	// (GET /api/v1/issuers/{id})
	GetIssuerByID(ctx context.Context, request GetIssuerByIDRequestObject) (GetIssuerByIDResponseObject, error)
	// Updates an issuer.
	// (PATCH /api/v1/issuers/{id})
	UpdateIssuer(ctx context.Context, request UpdateIssuerRequestObject) (UpdateIssuerResponseObject, error)
	// Creates an OAuth client.
	// (POST /api/v1/owners/{ownerID}/clients)
	CreateOAuthClient(ctx context.Context, request CreateOAuthClientRequestObject) (CreateOAuthClientResponseObject, error)
	// Creates an issuer.
	// (POST /api/v1/owners/{ownerID}/issuers)
	CreateIssuer(ctx context.Context, request CreateIssuerRequestObject) (CreateIssuerResponseObject, error)
	// Gets information about a User.
	// (GET /api/v1/users/{userID})
	GetUserByID(ctx context.Context, request GetUserByIDRequestObject) (GetUserByIDResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// DeleteOAuthClient operation middleware
func (sh *strictHandler) DeleteOAuthClient(ctx echo.Context, clientID gidx.PrefixedID) error {
	var request DeleteOAuthClientRequestObject

	request.ClientID = clientID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteOAuthClient(ctx.Request().Context(), request.(DeleteOAuthClientRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteOAuthClient")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteOAuthClientResponseObject); ok {
		return validResponse.VisitDeleteOAuthClientResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetOAuthClient operation middleware
func (sh *strictHandler) GetOAuthClient(ctx echo.Context, clientID gidx.PrefixedID) error {
	var request GetOAuthClientRequestObject

	request.ClientID = clientID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetOAuthClient(ctx.Request().Context(), request.(GetOAuthClientRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetOAuthClient")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetOAuthClientResponseObject); ok {
		return validResponse.VisitGetOAuthClientResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteIssuer operation middleware
func (sh *strictHandler) DeleteIssuer(ctx echo.Context, id gidx.PrefixedID) error {
	var request DeleteIssuerRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteIssuer(ctx.Request().Context(), request.(DeleteIssuerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteIssuer")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteIssuerResponseObject); ok {
		return validResponse.VisitDeleteIssuerResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetIssuerByID operation middleware
func (sh *strictHandler) GetIssuerByID(ctx echo.Context, id gidx.PrefixedID) error {
	var request GetIssuerByIDRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetIssuerByID(ctx.Request().Context(), request.(GetIssuerByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetIssuerByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetIssuerByIDResponseObject); ok {
		return validResponse.VisitGetIssuerByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// UpdateIssuer operation middleware
func (sh *strictHandler) UpdateIssuer(ctx echo.Context, id gidx.PrefixedID) error {
	var request UpdateIssuerRequestObject

	request.Id = id

	var body UpdateIssuerJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateIssuer(ctx.Request().Context(), request.(UpdateIssuerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateIssuer")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(UpdateIssuerResponseObject); ok {
		return validResponse.VisitUpdateIssuerResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// CreateOAuthClient operation middleware
func (sh *strictHandler) CreateOAuthClient(ctx echo.Context, ownerID gidx.PrefixedID) error {
	var request CreateOAuthClientRequestObject

	request.OwnerID = ownerID

	var body CreateOAuthClientJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateOAuthClient(ctx.Request().Context(), request.(CreateOAuthClientRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateOAuthClient")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(CreateOAuthClientResponseObject); ok {
		return validResponse.VisitCreateOAuthClientResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// CreateIssuer operation middleware
func (sh *strictHandler) CreateIssuer(ctx echo.Context, ownerID gidx.PrefixedID) error {
	var request CreateIssuerRequestObject

	request.OwnerID = ownerID

	var body CreateIssuerJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateIssuer(ctx.Request().Context(), request.(CreateIssuerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateIssuer")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(CreateIssuerResponseObject); ok {
		return validResponse.VisitCreateIssuerResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUserByID operation middleware
func (sh *strictHandler) GetUserByID(ctx echo.Context, userID gidx.PrefixedID) error {
	var request GetUserByIDRequestObject

	request.UserID = userID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUserByID(ctx.Request().Context(), request.(GetUserByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUserByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUserByIDResponseObject); ok {
		return validResponse.VisitGetUserByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
