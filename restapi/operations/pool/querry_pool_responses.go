// Code generated by go-swagger; DO NOT EDIT.

package pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"tam.io/homework/models"
)

// QuerryPoolOKCode is the HTTP code returned for type QuerryPoolOK
const QuerryPoolOKCode int = 200

/*QuerryPoolOK OK

swagger:response querryPoolOK
*/
type QuerryPoolOK struct {

	/*
	  In: Body
	*/
	Payload *models.PoolQueryResponse `json:"body,omitempty"`
}

// NewQuerryPoolOK creates QuerryPoolOK with default headers values
func NewQuerryPoolOK() *QuerryPoolOK {

	return &QuerryPoolOK{}
}

// WithPayload adds the payload to the querry pool o k response
func (o *QuerryPoolOK) WithPayload(payload *models.PoolQueryResponse) *QuerryPoolOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the querry pool o k response
func (o *QuerryPoolOK) SetPayload(payload *models.PoolQueryResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuerryPoolOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*QuerryPoolDefault generic error response

swagger:response querryPoolDefault
*/
type QuerryPoolDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload models.ErrorResponse `json:"body,omitempty"`
}

// NewQuerryPoolDefault creates QuerryPoolDefault with default headers values
func NewQuerryPoolDefault(code int) *QuerryPoolDefault {
	if code <= 0 {
		code = 500
	}

	return &QuerryPoolDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the querry pool default response
func (o *QuerryPoolDefault) WithStatusCode(code int) *QuerryPoolDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the querry pool default response
func (o *QuerryPoolDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the querry pool default response
func (o *QuerryPoolDefault) WithPayload(payload models.ErrorResponse) *QuerryPoolDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the querry pool default response
func (o *QuerryPoolDefault) SetPayload(payload models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuerryPoolDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
