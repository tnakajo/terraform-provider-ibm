// Code generated by go-swagger; DO NOT EDIT.

package l_baas

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

// NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams creates a new GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams object
// with the default values initialized.
func NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams() *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	var ()
	return &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParamsWithTimeout creates a new GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParamsWithTimeout(timeout time.Duration) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	var ()
	return &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams{

		timeout: timeout,
	}
}

// NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParamsWithContext creates a new GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParamsWithContext(ctx context.Context) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	var ()
	return &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams{

		Context: ctx,
	}
}

// NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParamsWithHTTPClient creates a new GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParamsWithHTTPClient(client *http.Client) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	var ()
	return &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams{
		HTTPClient: client,
	}
}

/*GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams contains all the parameters to send to the API endpoint
for the get load balancers ID listeners listener ID policies policy ID operation typically these are written to a http.Request
*/
type GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The load balancer identifier

	*/
	ID string
	/*ListenerID
	  The listener identifier

	*/
	ListenerID string
	/*PolicyID
	  The policy identifier

	*/
	PolicyID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithTimeout(timeout time.Duration) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithContext(ctx context.Context) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithHTTPClient(client *http.Client) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithGeneration(generation int64) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithID(id string) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetID(id string) {
	o.ID = id
}

// WithListenerID adds the listenerID to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithListenerID(listenerID string) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetListenerID(listenerID)
	return o
}

// SetListenerID adds the listenerId to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetListenerID(listenerID string) {
	o.ListenerID = listenerID
}

// WithPolicyID adds the policyID to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithPolicyID(policyID string) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetPolicyID(policyID string) {
	o.PolicyID = policyID
}

// WithVersion adds the version to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WithVersion(version string) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get load balancers ID listeners listener ID policies policy ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param listener_id
	if err := r.SetPathParam("listener_id", o.ListenerID); err != nil {
		return err
	}

	// path param policy_id
	if err := r.SetPathParam("policy_id", o.PolicyID); err != nil {
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
