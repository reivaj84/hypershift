// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudIpsecpoliciesPostReader is a Reader for the PcloudIpsecpoliciesPost structure.
type PcloudIpsecpoliciesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudIpsecpoliciesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudIpsecpoliciesPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudIpsecpoliciesPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudIpsecpoliciesPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudIpsecpoliciesPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudIpsecpoliciesPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudIpsecpoliciesPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudIpsecpoliciesPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudIpsecpoliciesPostOK creates a PcloudIpsecpoliciesPostOK with default headers values
func NewPcloudIpsecpoliciesPostOK() *PcloudIpsecpoliciesPostOK {
	return &PcloudIpsecpoliciesPostOK{}
}

/* PcloudIpsecpoliciesPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudIpsecpoliciesPostOK struct {
	Payload *models.IPSecPolicy
}

func (o *PcloudIpsecpoliciesPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies][%d] pcloudIpsecpoliciesPostOK  %+v", 200, o.Payload)
}
func (o *PcloudIpsecpoliciesPostOK) GetPayload() *models.IPSecPolicy {
	return o.Payload
}

func (o *PcloudIpsecpoliciesPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPSecPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesPostBadRequest creates a PcloudIpsecpoliciesPostBadRequest with default headers values
func NewPcloudIpsecpoliciesPostBadRequest() *PcloudIpsecpoliciesPostBadRequest {
	return &PcloudIpsecpoliciesPostBadRequest{}
}

/* PcloudIpsecpoliciesPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudIpsecpoliciesPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies][%d] pcloudIpsecpoliciesPostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudIpsecpoliciesPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIpsecpoliciesPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesPostUnauthorized creates a PcloudIpsecpoliciesPostUnauthorized with default headers values
func NewPcloudIpsecpoliciesPostUnauthorized() *PcloudIpsecpoliciesPostUnauthorized {
	return &PcloudIpsecpoliciesPostUnauthorized{}
}

/* PcloudIpsecpoliciesPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudIpsecpoliciesPostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies][%d] pcloudIpsecpoliciesPostUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudIpsecpoliciesPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIpsecpoliciesPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesPostForbidden creates a PcloudIpsecpoliciesPostForbidden with default headers values
func NewPcloudIpsecpoliciesPostForbidden() *PcloudIpsecpoliciesPostForbidden {
	return &PcloudIpsecpoliciesPostForbidden{}
}

/* PcloudIpsecpoliciesPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudIpsecpoliciesPostForbidden struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies][%d] pcloudIpsecpoliciesPostForbidden  %+v", 403, o.Payload)
}
func (o *PcloudIpsecpoliciesPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIpsecpoliciesPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesPostConflict creates a PcloudIpsecpoliciesPostConflict with default headers values
func NewPcloudIpsecpoliciesPostConflict() *PcloudIpsecpoliciesPostConflict {
	return &PcloudIpsecpoliciesPostConflict{}
}

/* PcloudIpsecpoliciesPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudIpsecpoliciesPostConflict struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies][%d] pcloudIpsecpoliciesPostConflict  %+v", 409, o.Payload)
}
func (o *PcloudIpsecpoliciesPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIpsecpoliciesPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesPostUnprocessableEntity creates a PcloudIpsecpoliciesPostUnprocessableEntity with default headers values
func NewPcloudIpsecpoliciesPostUnprocessableEntity() *PcloudIpsecpoliciesPostUnprocessableEntity {
	return &PcloudIpsecpoliciesPostUnprocessableEntity{}
}

/* PcloudIpsecpoliciesPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudIpsecpoliciesPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies][%d] pcloudIpsecpoliciesPostUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudIpsecpoliciesPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIpsecpoliciesPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesPostInternalServerError creates a PcloudIpsecpoliciesPostInternalServerError with default headers values
func NewPcloudIpsecpoliciesPostInternalServerError() *PcloudIpsecpoliciesPostInternalServerError {
	return &PcloudIpsecpoliciesPostInternalServerError{}
}

/* PcloudIpsecpoliciesPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudIpsecpoliciesPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ipsec-policies][%d] pcloudIpsecpoliciesPostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudIpsecpoliciesPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIpsecpoliciesPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}