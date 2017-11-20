// Code generated by go-swagger; DO NOT EDIT.

package account_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/alien45/bridge-api/models"
)

// GetPortfoliosByAccountOKCode is the HTTP code returned for type GetPortfoliosByAccountOK
const GetPortfoliosByAccountOKCode int = 200

/*GetPortfoliosByAccountOK successful operation

swagger:response getPortfoliosByAccountOK
*/
type GetPortfoliosByAccountOK struct {

	/*
	  In: Body
	*/
	Payload models.GetPortfoliosByAccountOKBody `json:"body,omitempty"`
}

// NewGetPortfoliosByAccountOK creates GetPortfoliosByAccountOK with default headers values
func NewGetPortfoliosByAccountOK() *GetPortfoliosByAccountOK {
	return &GetPortfoliosByAccountOK{}
}

// WithPayload adds the payload to the get portfolios by account o k response
func (o *GetPortfoliosByAccountOK) WithPayload(payload models.GetPortfoliosByAccountOKBody) *GetPortfoliosByAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get portfolios by account o k response
func (o *GetPortfoliosByAccountOK) SetPayload(payload models.GetPortfoliosByAccountOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPortfoliosByAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetPortfoliosByAccountOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetPortfoliosByAccountDefault unexpected error

swagger:response getPortfoliosByAccountDefault
*/
type GetPortfoliosByAccountDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetPortfoliosByAccountDefault creates GetPortfoliosByAccountDefault with default headers values
func NewGetPortfoliosByAccountDefault(code int) *GetPortfoliosByAccountDefault {
	if code <= 0 {
		code = 500
	}

	return &GetPortfoliosByAccountDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get portfolios by account default response
func (o *GetPortfoliosByAccountDefault) WithStatusCode(code int) *GetPortfoliosByAccountDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get portfolios by account default response
func (o *GetPortfoliosByAccountDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get portfolios by account default response
func (o *GetPortfoliosByAccountDefault) WithPayload(payload *models.ErrorResponse) *GetPortfoliosByAccountDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get portfolios by account default response
func (o *GetPortfoliosByAccountDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPortfoliosByAccountDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}