// Code generated by go-swagger; DO NOT EDIT.

package pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"tam.io/homework/models"
)

// NewQuerryPoolParams creates a new QuerryPoolParams object
// no default values defined in spec.
func NewQuerryPoolParams() QuerryPoolParams {

	return QuerryPoolParams{}
}

// QuerryPoolParams contains all the bound params for the querry pool operation
// typically these are obtained from a http.Request
//
// swagger:parameters querryPool
type QuerryPoolParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body *models.PoolQueryRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewQuerryPoolParams() beforehand.
func (o *QuerryPoolParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PoolQueryRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}