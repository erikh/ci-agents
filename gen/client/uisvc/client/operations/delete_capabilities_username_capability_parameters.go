// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteCapabilitiesUsernameCapabilityParams creates a new DeleteCapabilitiesUsernameCapabilityParams object
// with the default values initialized.
func NewDeleteCapabilitiesUsernameCapabilityParams() *DeleteCapabilitiesUsernameCapabilityParams {
	var ()
	return &DeleteCapabilitiesUsernameCapabilityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCapabilitiesUsernameCapabilityParamsWithTimeout creates a new DeleteCapabilitiesUsernameCapabilityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCapabilitiesUsernameCapabilityParamsWithTimeout(timeout time.Duration) *DeleteCapabilitiesUsernameCapabilityParams {
	var ()
	return &DeleteCapabilitiesUsernameCapabilityParams{

		timeout: timeout,
	}
}

// NewDeleteCapabilitiesUsernameCapabilityParamsWithContext creates a new DeleteCapabilitiesUsernameCapabilityParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCapabilitiesUsernameCapabilityParamsWithContext(ctx context.Context) *DeleteCapabilitiesUsernameCapabilityParams {
	var ()
	return &DeleteCapabilitiesUsernameCapabilityParams{

		Context: ctx,
	}
}

// NewDeleteCapabilitiesUsernameCapabilityParamsWithHTTPClient creates a new DeleteCapabilitiesUsernameCapabilityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCapabilitiesUsernameCapabilityParamsWithHTTPClient(client *http.Client) *DeleteCapabilitiesUsernameCapabilityParams {
	var ()
	return &DeleteCapabilitiesUsernameCapabilityParams{
		HTTPClient: client,
	}
}

/*DeleteCapabilitiesUsernameCapabilityParams contains all the parameters to send to the API endpoint
for the delete capabilities username capability operation typically these are written to a http.Request
*/
type DeleteCapabilitiesUsernameCapabilityParams struct {

	/*Capability
	  The name of the capability to remove.


	*/
	Capability string
	/*Username
	  The user ID to remove the capability from.


	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete capabilities username capability params
func (o *DeleteCapabilitiesUsernameCapabilityParams) WithTimeout(timeout time.Duration) *DeleteCapabilitiesUsernameCapabilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete capabilities username capability params
func (o *DeleteCapabilitiesUsernameCapabilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete capabilities username capability params
func (o *DeleteCapabilitiesUsernameCapabilityParams) WithContext(ctx context.Context) *DeleteCapabilitiesUsernameCapabilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete capabilities username capability params
func (o *DeleteCapabilitiesUsernameCapabilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete capabilities username capability params
func (o *DeleteCapabilitiesUsernameCapabilityParams) WithHTTPClient(client *http.Client) *DeleteCapabilitiesUsernameCapabilityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete capabilities username capability params
func (o *DeleteCapabilitiesUsernameCapabilityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCapability adds the capability to the delete capabilities username capability params
func (o *DeleteCapabilitiesUsernameCapabilityParams) WithCapability(capability string) *DeleteCapabilitiesUsernameCapabilityParams {
	o.SetCapability(capability)
	return o
}

// SetCapability adds the capability to the delete capabilities username capability params
func (o *DeleteCapabilitiesUsernameCapabilityParams) SetCapability(capability string) {
	o.Capability = capability
}

// WithUsername adds the username to the delete capabilities username capability params
func (o *DeleteCapabilitiesUsernameCapabilityParams) WithUsername(username string) *DeleteCapabilitiesUsernameCapabilityParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the delete capabilities username capability params
func (o *DeleteCapabilitiesUsernameCapabilityParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCapabilitiesUsernameCapabilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param capability
	if err := r.SetPathParam("capability", o.Capability); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
