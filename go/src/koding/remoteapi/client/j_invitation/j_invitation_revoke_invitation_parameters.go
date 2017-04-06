package j_invitation

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

// NewJInvitationRevokeInvitationParams creates a new JInvitationRevokeInvitationParams object
// with the default values initialized.
func NewJInvitationRevokeInvitationParams() *JInvitationRevokeInvitationParams {
	var ()
	return &JInvitationRevokeInvitationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJInvitationRevokeInvitationParamsWithTimeout creates a new JInvitationRevokeInvitationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJInvitationRevokeInvitationParamsWithTimeout(timeout time.Duration) *JInvitationRevokeInvitationParams {
	var ()
	return &JInvitationRevokeInvitationParams{

		timeout: timeout,
	}
}

// NewJInvitationRevokeInvitationParamsWithContext creates a new JInvitationRevokeInvitationParams object
// with the default values initialized, and the ability to set a context for a request
func NewJInvitationRevokeInvitationParamsWithContext(ctx context.Context) *JInvitationRevokeInvitationParams {
	var ()
	return &JInvitationRevokeInvitationParams{

		Context: ctx,
	}
}

/*JInvitationRevokeInvitationParams contains all the parameters to send to the API endpoint
for the j invitation revoke invitation operation typically these are written to a http.Request
*/
type JInvitationRevokeInvitationParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j invitation revoke invitation params
func (o *JInvitationRevokeInvitationParams) WithTimeout(timeout time.Duration) *JInvitationRevokeInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j invitation revoke invitation params
func (o *JInvitationRevokeInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j invitation revoke invitation params
func (o *JInvitationRevokeInvitationParams) WithContext(ctx context.Context) *JInvitationRevokeInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j invitation revoke invitation params
func (o *JInvitationRevokeInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j invitation revoke invitation params
func (o *JInvitationRevokeInvitationParams) WithBody(body models.DefaultSelector) *JInvitationRevokeInvitationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j invitation revoke invitation params
func (o *JInvitationRevokeInvitationParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JInvitationRevokeInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
