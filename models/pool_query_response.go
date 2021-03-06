// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PoolQueryResponse pool query response
//
// swagger:model poolQueryResponse
type PoolQueryResponse struct {

	// calculated quantile
	// Required: true
	CalculatedQuantile *int32 `json:"calculatedQuantile"`

	// total count
	// Required: true
	TotalCount *int64 `json:"totalCount"`
}

// Validate validates this pool query response
func (m *PoolQueryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCalculatedQuantile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolQueryResponse) validateCalculatedQuantile(formats strfmt.Registry) error {

	if err := validate.Required("calculatedQuantile", "body", m.CalculatedQuantile); err != nil {
		return err
	}

	return nil
}

func (m *PoolQueryResponse) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("totalCount", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pool query response based on context it is used
func (m *PoolQueryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PoolQueryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolQueryResponse) UnmarshalBinary(b []byte) error {
	var res PoolQueryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
