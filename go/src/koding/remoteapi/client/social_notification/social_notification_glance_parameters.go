package social_notification

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

// NewSocialNotificationGlanceParams creates a new SocialNotificationGlanceParams object
// with the default values initialized.
func NewSocialNotificationGlanceParams() *SocialNotificationGlanceParams {
	var ()
	return &SocialNotificationGlanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSocialNotificationGlanceParamsWithTimeout creates a new SocialNotificationGlanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSocialNotificationGlanceParamsWithTimeout(timeout time.Duration) *SocialNotificationGlanceParams {
	var ()
	return &SocialNotificationGlanceParams{

		timeout: timeout,
	}
}

// NewSocialNotificationGlanceParamsWithContext creates a new SocialNotificationGlanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewSocialNotificationGlanceParamsWithContext(ctx context.Context) *SocialNotificationGlanceParams {
	var ()
	return &SocialNotificationGlanceParams{

		Context: ctx,
	}
}

/*SocialNotificationGlanceParams contains all the parameters to send to the API endpoint
for the social notification glance operation typically these are written to a http.Request
*/
type SocialNotificationGlanceParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the social notification glance params
func (o *SocialNotificationGlanceParams) WithTimeout(timeout time.Duration) *SocialNotificationGlanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the social notification glance params
func (o *SocialNotificationGlanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the social notification glance params
func (o *SocialNotificationGlanceParams) WithContext(ctx context.Context) *SocialNotificationGlanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the social notification glance params
func (o *SocialNotificationGlanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the social notification glance params
func (o *SocialNotificationGlanceParams) WithBody(body models.DefaultSelector) *SocialNotificationGlanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the social notification glance params
func (o *SocialNotificationGlanceParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SocialNotificationGlanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
