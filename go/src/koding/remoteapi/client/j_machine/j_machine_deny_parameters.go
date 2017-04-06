package j_machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewJMachineDenyParams creates a new JMachineDenyParams object
// with the default values initialized.
func NewJMachineDenyParams() *JMachineDenyParams {
	var ()
	return &JMachineDenyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJMachineDenyParamsWithTimeout creates a new JMachineDenyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJMachineDenyParamsWithTimeout(timeout time.Duration) *JMachineDenyParams {
	var ()
	return &JMachineDenyParams{

		timeout: timeout,
	}
}

// NewJMachineDenyParamsWithContext creates a new JMachineDenyParams object
// with the default values initialized, and the ability to set a context for a request
func NewJMachineDenyParamsWithContext(ctx context.Context) *JMachineDenyParams {
	var ()
	return &JMachineDenyParams{

		Context: ctx,
	}
}

/*JMachineDenyParams contains all the parameters to send to the API endpoint
for the j machine deny operation typically these are written to a http.Request
*/
type JMachineDenyParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j machine deny params
func (o *JMachineDenyParams) WithTimeout(timeout time.Duration) *JMachineDenyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j machine deny params
func (o *JMachineDenyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j machine deny params
func (o *JMachineDenyParams) WithContext(ctx context.Context) *JMachineDenyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j machine deny params
func (o *JMachineDenyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j machine deny params
func (o *JMachineDenyParams) WithBody(body models.DefaultSelector) *JMachineDenyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j machine deny params
func (o *JMachineDenyParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j machine deny params
func (o *JMachineDenyParams) WithID(id string) *JMachineDenyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j machine deny params
func (o *JMachineDenyParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JMachineDenyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
