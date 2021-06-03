// Code generated by go-swagger; DO NOT EDIT.

package refplant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetRefPlantsURL generates an URL for the get ref plants operation
type GetRefPlantsURL struct {
	Category          *int32
	Classifiers       *string
	FloweringTime     *string
	Hight             *string
	Kind              *string
	Limit             int32
	Offset            int32
	RecommendPosition *string
	RegardToLight     *string
	RegardToMoisture  *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetRefPlantsURL) WithBasePath(bp string) *GetRefPlantsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetRefPlantsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetRefPlantsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/api/v1/refplants"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var categoryQ string
	if o.Category != nil {
		categoryQ = swag.FormatInt32(*o.Category)
	}
	if categoryQ != "" {
		qs.Set("category", categoryQ)
	}

	var classifiersQ string
	if o.Classifiers != nil {
		classifiersQ = *o.Classifiers
	}
	if classifiersQ != "" {
		qs.Set("classifiers", classifiersQ)
	}

	var floweringTimeQ string
	if o.FloweringTime != nil {
		floweringTimeQ = *o.FloweringTime
	}
	if floweringTimeQ != "" {
		qs.Set("floweringTime", floweringTimeQ)
	}

	var hightQ string
	if o.Hight != nil {
		hightQ = *o.Hight
	}
	if hightQ != "" {
		qs.Set("hight", hightQ)
	}

	var kindQ string
	if o.Kind != nil {
		kindQ = *o.Kind
	}
	if kindQ != "" {
		qs.Set("kind", kindQ)
	}

	limitQ := swag.FormatInt32(o.Limit)
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	offsetQ := swag.FormatInt32(o.Offset)
	if offsetQ != "" {
		qs.Set("offset", offsetQ)
	}

	var recommendPositionQ string
	if o.RecommendPosition != nil {
		recommendPositionQ = *o.RecommendPosition
	}
	if recommendPositionQ != "" {
		qs.Set("recommendPosition", recommendPositionQ)
	}

	var regardToLightQ string
	if o.RegardToLight != nil {
		regardToLightQ = *o.RegardToLight
	}
	if regardToLightQ != "" {
		qs.Set("regardToLight", regardToLightQ)
	}

	var regardToMoistureQ string
	if o.RegardToMoisture != nil {
		regardToMoistureQ = *o.RegardToMoisture
	}
	if regardToMoistureQ != "" {
		qs.Set("regardToMoisture", regardToMoistureQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetRefPlantsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetRefPlantsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetRefPlantsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetRefPlantsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetRefPlantsURL")
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
func (o *GetRefPlantsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
