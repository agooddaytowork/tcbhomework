// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ErrorResponse error response
//
// swagger:model errorResponse
type ErrorResponse string

// Validate validates this error response
func (m ErrorResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this error response based on context it is used
func (m ErrorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
