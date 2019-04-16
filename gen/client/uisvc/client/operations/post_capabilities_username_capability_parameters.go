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

// NewPostCapabilitiesUsernameCapabilityParams creates a new PostCapabilitiesUsernameCapabilityParams object
// with the default values initialized.
func NewPostCapabilitiesUsernameCapabilityParams() *PostCapabilitiesUsernameCapabilityParams {
	var ()
	return &PostCapabilitiesUsernameCapabilityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCapabilitiesUsernameCapabilityParamsWithTimeout creates a new PostCapabilitiesUsernameCapabilityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCapabilitiesUsernameCapabilityParamsWithTimeout(timeout time.Duration) *PostCapabilitiesUsernameCapabilityParams {
	var ()
	return &PostCapabilitiesUsernameCapabilityParams{

		timeout: timeout,
	}
}

// NewPostCapabilitiesUsernameCapabilityParamsWithContext creates a new PostCapabilitiesUsernameCapabilityParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCapabilitiesUsernameCapabilityParamsWithContext(ctx context.Context) *PostCapabilitiesUsernameCapabilityParams {
	var ()
	return &PostCapabilitiesUsernameCapabilityParams{

		Context: ctx,
	}
}

// NewPostCapabilitiesUsernameCapabilityParamsWithHTTPClient creates a new PostCapabilitiesUsernameCapabilityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCapabilitiesUsernameCapabilityParamsWithHTTPClient(client *http.Client) *PostCapabilitiesUsernameCapabilityParams {
	var ()
	return &PostCapabilitiesUsernameCapabilityParams{
		HTTPClient: client,
	}
}

/*PostCapabilitiesUsernameCapabilityParams contains all the parameters to send to the API endpoint
for the post capabilities username capability operation typically these are written to a http.Request
*/
type PostCapabilitiesUsernameCapabilityParams struct {

	/*Capability
	  The name of the capability to add.


	*/
	Capability string
	/*Username
	  The user ID to add the capability to.


	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post capabilities username capability params
func (o *PostCapabilitiesUsernameCapabilityParams) WithTimeout(timeout time.Duration) *PostCapabilitiesUsernameCapabilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post capabilities username capability params
func (o *PostCapabilitiesUsernameCapabilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post capabilities username capability params
func (o *PostCapabilitiesUsernameCapabilityParams) WithContext(ctx context.Context) *PostCapabilitiesUsernameCapabilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post capabilities username capability params
func (o *PostCapabilitiesUsernameCapabilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post capabilities username capability params
func (o *PostCapabilitiesUsernameCapabilityParams) WithHTTPClient(client *http.Client) *PostCapabilitiesUsernameCapabilityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post capabilities username capability params
func (o *PostCapabilitiesUsernameCapabilityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCapability adds the capability to the post capabilities username capability params
func (o *PostCapabilitiesUsernameCapabilityParams) WithCapability(capability string) *PostCapabilitiesUsernameCapabilityParams {
	o.SetCapability(capability)
	return o
}

// SetCapability adds the capability to the post capabilities username capability params
func (o *PostCapabilitiesUsernameCapabilityParams) SetCapability(capability string) {
	o.Capability = capability
}

// WithUsername adds the username to the post capabilities username capability params
func (o *PostCapabilitiesUsernameCapabilityParams) WithUsername(username string) *PostCapabilitiesUsernameCapabilityParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the post capabilities username capability params
func (o *PostCapabilitiesUsernameCapabilityParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *PostCapabilitiesUsernameCapabilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
