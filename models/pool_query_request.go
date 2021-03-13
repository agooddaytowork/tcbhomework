// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PoolQueryRequest pool query request
//
// swagger:model poolQueryRequest
type PoolQueryRequest struct {

	// percentile
	// Required: true
	Percentile *float64 `json:"percentile"`

	// pool Id
	// Required: true
	PoolID *int64 `json:"poolId"`
}

// Validate validates this pool query request
func (m *PoolQueryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePercentile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoolID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolQueryRequest) validatePercentile(formats strfmt.Registry) error {

	if err := validate.Required("percentile", "body", m.Percentile); err != nil {
		return err
	}

	return nil
}

func (m *PoolQueryRequest) validatePoolID(formats strfmt.Registry) error {

	if err := validate.Required("poolId", "body", m.PoolID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolQueryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolQueryRequest) UnmarshalBinary(b []byte) error {
	var res PoolQueryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
