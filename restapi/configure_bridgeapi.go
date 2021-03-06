// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/alien45/bridge-api/restapi/operations"
	"github.com/alien45/bridge-api/restapi/operations/account_data"
	"github.com/alien45/bridge-api/restapi/operations/market_access"
	"github.com/alien45/bridge-api/restapi/operations/system_info"
	"github.com/alien45/bridge-api/models"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name bridgeapi --spec ../api/swagger.yaml

func configureFlags(api *operations.BridgeapiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BridgeapiAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BridgeAuthAuth = func(token string, scopes []string) (interface{}, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (bridge_auth) has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	api.SystemInfoSystemInfoHandler = system_info.SystemInfoHandlerFunc(func(params system_info.SystemInfoParams) middleware.Responder {
		return middleware.NotImplemented("operation system_info.SystemInfo has not yet been implemented")
	})
	api.MarketAccessAddOrderHandler = market_access.AddOrderHandlerFunc(func(params market_access.AddOrderParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation market_access.AddOrder has not yet been implemented")
	})
	api.MarketAccessCancelOrderHandler = market_access.CancelOrderHandlerFunc(func(params market_access.CancelOrderParams) middleware.Responder {
		return middleware.NotImplemented("operation market_access.CancelOrder has not yet been implemented")
	})
	api.MarketAccessGetExchangesHandler = market_access.GetExchangesHandlerFunc(func(params market_access.GetExchangesParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation market_access.GetExchanges has not yet been implemented")
	})
	api.AccountDataGetLinkedAccountsHandler = account_data.GetLinkedAccountsHandlerFunc(func(params account_data.GetLinkedAccountsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation account_data.GetLinkedAccounts has not yet been implemented")
	})
	api.MarketAccessGetOrderByIDHandler = market_access.GetOrderByIDHandlerFunc(func(params market_access.GetOrderByIDParams) middleware.Responder {
		//return middleware.NotImplemented("operation market_access.GetOrderByID has not yet been implemented")
		// ToDo: Remove sample before deploying
		sampleOrder := models.Order{
			ID: params.OrderID,
			Quantity: 99.00,
			Rate: 1.99,
		}
		return market_access.NewGetOrderByIDOK().WithPayload(&sampleOrder);
	})
	api.AccountDataGetOrdersByAccountHandler = account_data.GetOrdersByAccountHandlerFunc(func(params account_data.GetOrdersByAccountParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation account_data.GetOrdersByAccount has not yet been implemented")
	})
	api.AccountDataGetPortfoliosByAccountHandler = account_data.GetPortfoliosByAccountHandlerFunc(func(params account_data.GetPortfoliosByAccountParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation account_data.GetPortfoliosByAccount has not yet been implemented")
	})
	api.MarketAccessGetSymbolsHandler = market_access.GetSymbolsHandlerFunc(func(params market_access.GetSymbolsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation market_access.GetSymbols has not yet been implemented")
	})
	api.MarketAccessGetTickerHandler = market_access.GetTickerHandlerFunc(func(params market_access.GetTickerParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation market_access.GetTicker has not yet been implemented")
	})
	api.AccountDataGetTradesByAccountHandler = account_data.GetTradesByAccountHandlerFunc(func(params account_data.GetTradesByAccountParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation account_data.GetTradesByAccount has not yet been implemented")
	})
	api.MarketAccessUpdateOrderHandler = market_access.UpdateOrderHandlerFunc(func(params market_access.UpdateOrderParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation market_access.UpdateOrder has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
