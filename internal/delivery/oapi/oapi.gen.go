// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package oapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Item defines model for Item.
type Item struct {
	// Price The total price payed for this item.
	Price string `json:"price"`

	// ShortDescription The Short Product Description for the item.
	ShortDescription string `json:"shortDescription"`
}

// Receipt defines model for Receipt.
type Receipt struct {
	Items []Item `json:"items"`

	// PurchaseDate The date of the purchase printed on the receipt.
	PurchaseDate openapi_types.Date `json:"purchaseDate"`

	// PurchaseTime The time of the purchase printed on the receipt. 24-hour time expected.
	PurchaseTime string `json:"purchaseTime"`

	// Retailer The name of the retailer or store the receipt is from.
	Retailer string `json:"retailer"`

	// Total The total amount paid on the receipt.
	Total string `json:"total"`
}

// PostReceiptsProcessJSONRequestBody defines body for PostReceiptsProcess for application/json ContentType.
type PostReceiptsProcessJSONRequestBody = Receipt

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Submits a receipt for processing.
	// (POST /receipts/process)
	PostReceiptsProcess(w http.ResponseWriter, r *http.Request)
	// Deletes a receipt
	// (DELETE /receipts/{id})
	DeleteReceiptsId(w http.ResponseWriter, r *http.Request, id string)
	// Returns the points awarded for the receipt.
	// (GET /receipts/{id}/points)
	GetReceiptsIdPoints(w http.ResponseWriter, r *http.Request, id string)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Submits a receipt for processing.
// (POST /receipts/process)
func (_ Unimplemented) PostReceiptsProcess(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Deletes a receipt
// (DELETE /receipts/{id})
func (_ Unimplemented) DeleteReceiptsId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Returns the points awarded for the receipt.
// (GET /receipts/{id}/points)
func (_ Unimplemented) GetReceiptsIdPoints(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostReceiptsProcess operation middleware
func (siw *ServerInterfaceWrapper) PostReceiptsProcess(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostReceiptsProcess(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteReceiptsId operation middleware
func (siw *ServerInterfaceWrapper) DeleteReceiptsId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteReceiptsId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetReceiptsIdPoints operation middleware
func (siw *ServerInterfaceWrapper) GetReceiptsIdPoints(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetReceiptsIdPoints(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/receipts/process", wrapper.PostReceiptsProcess)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/receipts/{id}", wrapper.DeleteReceiptsId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/receipts/{id}/points", wrapper.GetReceiptsIdPoints)
	})

	return r
}

type BadRequestResponse struct {
}

type NotFoundResponse struct {
}

type PostReceiptsProcessRequestObject struct {
	Body *PostReceiptsProcessJSONRequestBody
}

type PostReceiptsProcessResponseObject interface {
	VisitPostReceiptsProcessResponse(w http.ResponseWriter) error
}

type PostReceiptsProcess200JSONResponse struct {
	Id string `json:"id"`
}

func (response PostReceiptsProcess200JSONResponse) VisitPostReceiptsProcessResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostReceiptsProcess400Response = BadRequestResponse

func (response PostReceiptsProcess400Response) VisitPostReceiptsProcessResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type DeleteReceiptsIdRequestObject struct {
	Id string `json:"id"`
}

type DeleteReceiptsIdResponseObject interface {
	VisitDeleteReceiptsIdResponse(w http.ResponseWriter) error
}

type DeleteReceiptsId204Response struct {
}

func (response DeleteReceiptsId204Response) VisitDeleteReceiptsIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteReceiptsId404Response = NotFoundResponse

func (response DeleteReceiptsId404Response) VisitDeleteReceiptsIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetReceiptsIdPointsRequestObject struct {
	Id string `json:"id"`
}

type GetReceiptsIdPointsResponseObject interface {
	VisitGetReceiptsIdPointsResponse(w http.ResponseWriter) error
}

type GetReceiptsIdPoints200JSONResponse struct {
	Points *int64 `json:"points,omitempty"`
}

func (response GetReceiptsIdPoints200JSONResponse) VisitGetReceiptsIdPointsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetReceiptsIdPoints404Response = NotFoundResponse

func (response GetReceiptsIdPoints404Response) VisitGetReceiptsIdPointsResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Submits a receipt for processing.
	// (POST /receipts/process)
	PostReceiptsProcess(ctx context.Context, request PostReceiptsProcessRequestObject) (PostReceiptsProcessResponseObject, error)
	// Deletes a receipt
	// (DELETE /receipts/{id})
	DeleteReceiptsId(ctx context.Context, request DeleteReceiptsIdRequestObject) (DeleteReceiptsIdResponseObject, error)
	// Returns the points awarded for the receipt.
	// (GET /receipts/{id}/points)
	GetReceiptsIdPoints(ctx context.Context, request GetReceiptsIdPointsRequestObject) (GetReceiptsIdPointsResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// PostReceiptsProcess operation middleware
func (sh *strictHandler) PostReceiptsProcess(w http.ResponseWriter, r *http.Request) {
	var request PostReceiptsProcessRequestObject

	var body PostReceiptsProcessJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostReceiptsProcess(ctx, request.(PostReceiptsProcessRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostReceiptsProcess")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostReceiptsProcessResponseObject); ok {
		if err := validResponse.VisitPostReceiptsProcessResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteReceiptsId operation middleware
func (sh *strictHandler) DeleteReceiptsId(w http.ResponseWriter, r *http.Request, id string) {
	var request DeleteReceiptsIdRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteReceiptsId(ctx, request.(DeleteReceiptsIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteReceiptsId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteReceiptsIdResponseObject); ok {
		if err := validResponse.VisitDeleteReceiptsIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetReceiptsIdPoints operation middleware
func (sh *strictHandler) GetReceiptsIdPoints(w http.ResponseWriter, r *http.Request, id string) {
	var request GetReceiptsIdPointsRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetReceiptsIdPoints(ctx, request.(GetReceiptsIdPointsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetReceiptsIdPoints")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetReceiptsIdPointsResponseObject); ok {
		if err := validResponse.VisitGetReceiptsIdPointsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xWX2/bNhD/KgTXt8qyrHhGo7d1xgZjSBEkfYsygBZPMTuLZI+npkag7z6QkmXLVpoE",
	"felTYulI/vj7c6cnXpjKGg2aHM+eOIKzRjsIPz4KeQNfa3Dkf0lwBSpLymie8c8bYAgFKEtMOab0N7FV",
	"MuZNxD8Z+svUWp4v+mT6NaWvYKVBRhtBbLWMedNE3BUbqEQ4fUVQ+b8WjQUk1WKyqAoYh0OGxJaFAmbF",
	"Dvbbe3gEVcwjDt9FZbfAM76I55c84lYQAfod/s1z+T7P4zyXT2nzjkecdtZXOkKlH/zF3MYgLY/PHYNx",
	"66vYNRpZF8SOyjs4MILmytSahNJsCY9sll7/M4R2l+ePee7yfHL/fgRZE3GEr7VCkDy7O4cZdazd9yvN",
	"+gsU5O900+pxTrQHOfznHULJM/7b9GCZaafXNIjVRLxSetXWz/rDBKLY+Ze2xmIjHCwFPSOhFATMlIGl",
	"fbVXVBNIZnR43jloSGCapOkkmU2SGY94abASxDPutxsTcr/1Z1U95yVVvRoIS+eTjamxXQTfLRQEcohv",
	"dpENofnaMWgIJNQWcByWFgdY+0pmkDkyCMegfCZLNKc2y+skSRdX7E+DGpBdCfwP6DmvtcWjjot4CNuP",
	"cigq72lmhfqxcm8P4onde8ZODHYic9QZeQ/9PAx+Y6VLc36rP5hTHm7PrkVTgHPGH0qKwkW6JPnk9+++",
	"Abp2i1mcxIknzljQwiqe8Ys4iS/aq29CwKbd9m7a7R9Sacba7229rhQ5Jo4aKu5hKf3gOfZpFr5+JXnG",
	"r42jDqLrIPKWSHD00cidP6QwmkCH84S1W1WE9dMvrm12bdpf6gX7ltIMlSKsITw4mjBpkrzp2JMWFUbM",
	"wUpCrhfr3xfJJAEoJ/N0XUwu5WwxkeX8Q3mRwIfLdXpqtdtXNFQln3HLUJMboBq1C1ZfLZlwTj1okIzM",
	"0P1NxOftvcdY7PmZHo3fMBrrqhK4e5X4vv7gpiclm9ZCW2g779Aay/B8b46VDByhqIAAHc/uxkK+Wh76",
	"UJ9q5d96O/OI+1bFM8/dqQmiY0Ff1uL+zDPz8zzsw+fqwpNQ1tvtjrX3lS3h85cJ7z9chnS39BzRPULv",
	"1BrVfUI9AI3hO3ijLWXiUaDsP1IGPA7l+RvooM11e84vrtDPpPrAZJ/sWZIcjU6laTE/wPAz+QEwiPJi",
	"RsMQrau1H5vliRLxzxrlLSI3TdP8HwAA//9SAjOBfQsAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
