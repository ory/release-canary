// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListOAuth2ClientsParams creates a new ListOAuth2ClientsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListOAuth2ClientsParams() *ListOAuth2ClientsParams {
	return &ListOAuth2ClientsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListOAuth2ClientsParamsWithTimeout creates a new ListOAuth2ClientsParams object
// with the ability to set a timeout on a request.
func NewListOAuth2ClientsParamsWithTimeout(timeout time.Duration) *ListOAuth2ClientsParams {
	return &ListOAuth2ClientsParams{
		timeout: timeout,
	}
}

// NewListOAuth2ClientsParamsWithContext creates a new ListOAuth2ClientsParams object
// with the ability to set a context for a request.
func NewListOAuth2ClientsParamsWithContext(ctx context.Context) *ListOAuth2ClientsParams {
	return &ListOAuth2ClientsParams{
		Context: ctx,
	}
}

// NewListOAuth2ClientsParamsWithHTTPClient creates a new ListOAuth2ClientsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListOAuth2ClientsParamsWithHTTPClient(client *http.Client) *ListOAuth2ClientsParams {
	return &ListOAuth2ClientsParams{
		HTTPClient: client,
	}
}

/* ListOAuth2ClientsParams contains all the parameters to send to the API endpoint
   for the list o auth2 clients operation.

   Typically these are written to a http.Request.
*/
type ListOAuth2ClientsParams struct {

	/* ClientName.

	   The name of the clients to filter by.
	*/
	ClientName *string

	/* Limit.

	   The maximum amount of clients to returned, upper bound is 500 clients.

	   Format: int64
	*/
	Limit *int64

	/* Offset.

	   The offset from where to start looking.

	   Format: int64
	*/
	Offset *int64

	/* Owner.

	   The owner of the clients to filter by.
	*/
	Owner *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list o auth2 clients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOAuth2ClientsParams) WithDefaults() *ListOAuth2ClientsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list o auth2 clients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOAuth2ClientsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) WithTimeout(timeout time.Duration) *ListOAuth2ClientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) WithContext(ctx context.Context) *ListOAuth2ClientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) WithHTTPClient(client *http.Client) *ListOAuth2ClientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientName adds the clientName to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) WithClientName(clientName *string) *ListOAuth2ClientsParams {
	o.SetClientName(clientName)
	return o
}

// SetClientName adds the clientName to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) SetClientName(clientName *string) {
	o.ClientName = clientName
}

// WithLimit adds the limit to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) WithLimit(limit *int64) *ListOAuth2ClientsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) WithOffset(offset *int64) *ListOAuth2ClientsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) WithOwner(owner *string) *ListOAuth2ClientsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the list o auth2 clients params
func (o *ListOAuth2ClientsParams) SetOwner(owner *string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *ListOAuth2ClientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientName != nil {

		// query param client_name
		var qrClientName string

		if o.ClientName != nil {
			qrClientName = *o.ClientName
		}
		qClientName := qrClientName
		if qClientName != "" {

			if err := r.SetQueryParam("client_name", qClientName); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Owner != nil {

		// query param owner
		var qrOwner string

		if o.Owner != nil {
			qrOwner = *o.Owner
		}
		qOwner := qrOwner
		if qOwner != "" {

			if err := r.SetQueryParam("owner", qOwner); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}