// Code generated by go-swagger; DO NOT EDIT.

package market_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSymbolsHandlerFunc turns a function with the right signature into a get symbols handler
type GetSymbolsHandlerFunc func(GetSymbolsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSymbolsHandlerFunc) Handle(params GetSymbolsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetSymbolsHandler interface for that can handle valid get symbols params
type GetSymbolsHandler interface {
	Handle(GetSymbolsParams, interface{}) middleware.Responder
}

// NewGetSymbols creates a new http.Handler for the get symbols operation
func NewGetSymbols(ctx *middleware.Context, handler GetSymbolsHandler) *GetSymbols {
	return &GetSymbols{Context: ctx, Handler: handler}
}

/*GetSymbols swagger:route GET /market/symbols Market Access getSymbols

Return the list of symbols for the given trading venue.

Return the list of symbols for the specified trading venue. If no trading venue (exchange_id) is specified, the default one will be used

*/
type GetSymbols struct {
	Context *middleware.Context
	Handler GetSymbolsHandler
}

func (o *GetSymbols) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSymbolsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
