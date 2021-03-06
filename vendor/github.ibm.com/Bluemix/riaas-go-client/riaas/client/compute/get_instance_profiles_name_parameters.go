// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInstanceProfilesNameParams creates a new GetInstanceProfilesNameParams object
// with the default values initialized.
func NewGetInstanceProfilesNameParams() *GetInstanceProfilesNameParams {
	var ()
	return &GetInstanceProfilesNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstanceProfilesNameParamsWithTimeout creates a new GetInstanceProfilesNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstanceProfilesNameParamsWithTimeout(timeout time.Duration) *GetInstanceProfilesNameParams {
	var ()
	return &GetInstanceProfilesNameParams{

		timeout: timeout,
	}
}

// NewGetInstanceProfilesNameParamsWithContext creates a new GetInstanceProfilesNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstanceProfilesNameParamsWithContext(ctx context.Context) *GetInstanceProfilesNameParams {
	var ()
	return &GetInstanceProfilesNameParams{

		Context: ctx,
	}
}

// NewGetInstanceProfilesNameParamsWithHTTPClient creates a new GetInstanceProfilesNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstanceProfilesNameParamsWithHTTPClient(client *http.Client) *GetInstanceProfilesNameParams {
	var ()
	return &GetInstanceProfilesNameParams{
		HTTPClient: client,
	}
}

/*GetInstanceProfilesNameParams contains all the parameters to send to the API endpoint
for the get instance profiles name operation typically these are written to a http.Request
*/
type GetInstanceProfilesNameParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Name
	  The instance profile name

	*/
	Name string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instance profiles name params
func (o *GetInstanceProfilesNameParams) WithTimeout(timeout time.Duration) *GetInstanceProfilesNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instance profiles name params
func (o *GetInstanceProfilesNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instance profiles name params
func (o *GetInstanceProfilesNameParams) WithContext(ctx context.Context) *GetInstanceProfilesNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instance profiles name params
func (o *GetInstanceProfilesNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instance profiles name params
func (o *GetInstanceProfilesNameParams) WithHTTPClient(client *http.Client) *GetInstanceProfilesNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instance profiles name params
func (o *GetInstanceProfilesNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get instance profiles name params
func (o *GetInstanceProfilesNameParams) WithGeneration(generation int64) *GetInstanceProfilesNameParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get instance profiles name params
func (o *GetInstanceProfilesNameParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithName adds the name to the get instance profiles name params
func (o *GetInstanceProfilesNameParams) WithName(name string) *GetInstanceProfilesNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get instance profiles name params
func (o *GetInstanceProfilesNameParams) SetName(name string) {
	o.Name = name
}

// WithVersion adds the version to the get instance profiles name params
func (o *GetInstanceProfilesNameParams) WithVersion(version string) *GetInstanceProfilesNameParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get instance profiles name params
func (o *GetInstanceProfilesNameParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstanceProfilesNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
