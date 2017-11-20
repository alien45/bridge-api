// Code generated by go-swagger; DO NOT EDIT.

package market_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/alien45/bridge-api/models"
)

// UpdateOrderOKCode is the HTTP code returned for type UpdateOrderOK
const UpdateOrderOKCode int = 200

/*UpdateOrderOK Order updated successfully

swagger:response updateOrderOK
*/
type UpdateOrderOK struct {
}

// NewUpdateOrderOK creates UpdateOrderOK with default headers values
func NewUpdateOrderOK() *UpdateOrderOK {
	return &UpdateOrderOK{}
}

// WriteResponse to the client
func (o *UpdateOrderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateOrderBadRequestCode is the HTTP code returned for type UpdateOrderBadRequest
const UpdateOrderBadRequestCode int = 400

/*UpdateOrderBadRequest Invalid request

swagger:response updateOrderBadRequest
*/
type UpdateOrderBadRequest struct {
}

// NewUpdateOrderBadRequest creates UpdateOrderBadRequest with default headers values
func NewUpdateOrderBadRequest() *UpdateOrderBadRequest {
	return &UpdateOrderBadRequest{}
}

// WriteResponse to the client
func (o *UpdateOrderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateOrderForbiddenCode is the HTTP code returned for type UpdateOrderForbidden
const UpdateOrderForbiddenCode int = 403

/*UpdateOrderForbidden Order already filled, can not be canceled or modified

swagger:response updateOrderForbidden
*/
type UpdateOrderForbidden struct {
}

// NewUpdateOrderForbidden creates UpdateOrderForbidden with default headers values
func NewUpdateOrderForbidden() *UpdateOrderForbidden {
	return &UpdateOrderForbidden{}
}

// WriteResponse to the client
func (o *UpdateOrderForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// UpdateOrderNotFoundCode is the HTTP code returned for type UpdateOrderNotFound
const UpdateOrderNotFoundCode int = 404

/*UpdateOrderNotFound Order not found

swagger:response updateOrderNotFound
*/
type UpdateOrderNotFound struct {
}

// NewUpdateOrderNotFound creates UpdateOrderNotFound with default headers values
func NewUpdateOrderNotFound() *UpdateOrderNotFound {
	return &UpdateOrderNotFound{}
}

// WriteResponse to the client
func (o *UpdateOrderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UpdateOrderUnprocessableEntityCode is the HTTP code returned for type UpdateOrderUnprocessableEntity
const UpdateOrderUnprocessableEntityCode int = 422

/*UpdateOrderUnprocessableEntity Validation exception

swagger:response updateOrderUnprocessableEntity
*/
type UpdateOrderUnprocessableEntity struct {
}

// NewUpdateOrderUnprocessableEntity creates UpdateOrderUnprocessableEntity with default headers values
func NewUpdateOrderUnprocessableEntity() *UpdateOrderUnprocessableEntity {
	return &UpdateOrderUnprocessableEntity{}
}

// WriteResponse to the client
func (o *UpdateOrderUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(422)
}

/*UpdateOrderDefault unexpected error

swagger:response updateOrderDefault
*/
type UpdateOrderDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateOrderDefault creates UpdateOrderDefault with default headers values
func NewUpdateOrderDefault(code int) *UpdateOrderDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateOrderDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update order default response
func (o *UpdateOrderDefault) WithStatusCode(code int) *UpdateOrderDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update order default response
func (o *UpdateOrderDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update order default response
func (o *UpdateOrderDefault) WithPayload(payload *models.ErrorResponse) *UpdateOrderDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update order default response
func (o *UpdateOrderDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrderDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}