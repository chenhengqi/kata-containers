// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kata-containers/kata-containers/src/runtime/virtcontainers/pkg/firecracker/client/models"
)

// GetFirecrackerVersionReader is a Reader for the GetFirecrackerVersion structure.
type GetFirecrackerVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFirecrackerVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFirecrackerVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFirecrackerVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFirecrackerVersionOK creates a GetFirecrackerVersionOK with default headers values
func NewGetFirecrackerVersionOK() *GetFirecrackerVersionOK {
	return &GetFirecrackerVersionOK{}
}

/* GetFirecrackerVersionOK describes a response with status code 200, with default header values.

OK
*/
type GetFirecrackerVersionOK struct {
	Payload *models.FirecrackerVersion
}

func (o *GetFirecrackerVersionOK) Error() string {
	return fmt.Sprintf("[GET /version][%d] getFirecrackerVersionOK  %+v", 200, o.Payload)
}
func (o *GetFirecrackerVersionOK) GetPayload() *models.FirecrackerVersion {
	return o.Payload
}

func (o *GetFirecrackerVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FirecrackerVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFirecrackerVersionDefault creates a GetFirecrackerVersionDefault with default headers values
func NewGetFirecrackerVersionDefault(code int) *GetFirecrackerVersionDefault {
	return &GetFirecrackerVersionDefault{
		_statusCode: code,
	}
}

/* GetFirecrackerVersionDefault describes a response with status code -1, with default header values.

Internal server error
*/
type GetFirecrackerVersionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get firecracker version default response
func (o *GetFirecrackerVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetFirecrackerVersionDefault) Error() string {
	return fmt.Sprintf("[GET /version][%d] getFirecrackerVersion default  %+v", o._statusCode, o.Payload)
}
func (o *GetFirecrackerVersionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFirecrackerVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
