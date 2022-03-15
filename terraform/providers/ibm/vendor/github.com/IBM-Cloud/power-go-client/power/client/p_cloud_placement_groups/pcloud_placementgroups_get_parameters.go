// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_placement_groups

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

// NewPcloudPlacementgroupsGetParams creates a new PcloudPlacementgroupsGetParams object
// with the default values initialized.
func NewPcloudPlacementgroupsGetParams() *PcloudPlacementgroupsGetParams {
	var ()
	return &PcloudPlacementgroupsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPlacementgroupsGetParamsWithTimeout creates a new PcloudPlacementgroupsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudPlacementgroupsGetParamsWithTimeout(timeout time.Duration) *PcloudPlacementgroupsGetParams {
	var ()
	return &PcloudPlacementgroupsGetParams{

		timeout: timeout,
	}
}

// NewPcloudPlacementgroupsGetParamsWithContext creates a new PcloudPlacementgroupsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudPlacementgroupsGetParamsWithContext(ctx context.Context) *PcloudPlacementgroupsGetParams {
	var ()
	return &PcloudPlacementgroupsGetParams{

		Context: ctx,
	}
}

// NewPcloudPlacementgroupsGetParamsWithHTTPClient creates a new PcloudPlacementgroupsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudPlacementgroupsGetParamsWithHTTPClient(client *http.Client) *PcloudPlacementgroupsGetParams {
	var ()
	return &PcloudPlacementgroupsGetParams{
		HTTPClient: client,
	}
}

/*PcloudPlacementgroupsGetParams contains all the parameters to send to the API endpoint
for the pcloud placementgroups get operation typically these are written to a http.Request
*/
type PcloudPlacementgroupsGetParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*PlacementGroupID
	  Placement Group ID

	*/
	PlacementGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud placementgroups get params
func (o *PcloudPlacementgroupsGetParams) WithTimeout(timeout time.Duration) *PcloudPlacementgroupsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud placementgroups get params
func (o *PcloudPlacementgroupsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud placementgroups get params
func (o *PcloudPlacementgroupsGetParams) WithContext(ctx context.Context) *PcloudPlacementgroupsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud placementgroups get params
func (o *PcloudPlacementgroupsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud placementgroups get params
func (o *PcloudPlacementgroupsGetParams) WithHTTPClient(client *http.Client) *PcloudPlacementgroupsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud placementgroups get params
func (o *PcloudPlacementgroupsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud placementgroups get params
func (o *PcloudPlacementgroupsGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPlacementgroupsGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud placementgroups get params
func (o *PcloudPlacementgroupsGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPlacementGroupID adds the placementGroupID to the pcloud placementgroups get params
func (o *PcloudPlacementgroupsGetParams) WithPlacementGroupID(placementGroupID string) *PcloudPlacementgroupsGetParams {
	o.SetPlacementGroupID(placementGroupID)
	return o
}

// SetPlacementGroupID adds the placementGroupId to the pcloud placementgroups get params
func (o *PcloudPlacementgroupsGetParams) SetPlacementGroupID(placementGroupID string) {
	o.PlacementGroupID = placementGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPlacementgroupsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param placement_group_id
	if err := r.SetPathParam("placement_group_id", o.PlacementGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}