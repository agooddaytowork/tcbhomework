// Code generated by go-swagger; DO NOT EDIT.

package pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// QuerryPoolURL generates an URL for the querry pool operation
type QuerryPoolURL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *QuerryPoolURL) WithBasePath(bp string) *QuerryPoolURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *QuerryPoolURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *QuerryPoolURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/api/pools/query"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *QuerryPoolURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *QuerryPoolURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *QuerryPoolURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on QuerryPoolURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on QuerryPoolURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *QuerryPoolURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}