// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetKeysReader is a Reader for the GetKeys structure.
type GetKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetKeysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKeysOK creates a GetKeysOK with default headers values
func NewGetKeysOK() *GetKeysOK {
	return &GetKeysOK{}
}

/*GetKeysOK handles this case with default header values.

dummy
*/
type GetKeysOK struct {
	Payload *models.GetKeysOKBody
}

func (o *GetKeysOK) Error() string {
	return fmt.Sprintf("[GET /keys][%d] getKeysOK  %+v", 200, o.Payload)
}

func (o *GetKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetKeysOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeysInternalServerError creates a GetKeysInternalServerError with default headers values
func NewGetKeysInternalServerError() *GetKeysInternalServerError {
	return &GetKeysInternalServerError{}
}

/*GetKeysInternalServerError handles this case with default header values.

error
*/
type GetKeysInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetKeysInternalServerError) Error() string {
	return fmt.Sprintf("[GET /keys][%d] getKeysInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKeysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
