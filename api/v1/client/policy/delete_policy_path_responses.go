package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/polyverse-security/cilium/api/v1/models"
)

// DeletePolicyPathReader is a Reader for the DeletePolicyPath structure.
type DeletePolicyPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePolicyPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeletePolicyPathNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeletePolicyPathInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeletePolicyPathNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeletePolicyPathFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePolicyPathNoContent creates a DeletePolicyPathNoContent with default headers values
func NewDeletePolicyPathNoContent() *DeletePolicyPathNoContent {
	return &DeletePolicyPathNoContent{}
}

/*DeletePolicyPathNoContent handles this case with default header values.

Success
*/
type DeletePolicyPathNoContent struct {
}

func (o *DeletePolicyPathNoContent) Error() string {
	return fmt.Sprintf("[DELETE /policy/{path}][%d] deletePolicyPathNoContent ", 204)
}

func (o *DeletePolicyPathNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePolicyPathInvalid creates a DeletePolicyPathInvalid with default headers values
func NewDeletePolicyPathInvalid() *DeletePolicyPathInvalid {
	return &DeletePolicyPathInvalid{}
}

/*DeletePolicyPathInvalid handles this case with default header values.

Invalid request
*/
type DeletePolicyPathInvalid struct {
	Payload models.Error
}

func (o *DeletePolicyPathInvalid) Error() string {
	return fmt.Sprintf("[DELETE /policy/{path}][%d] deletePolicyPathInvalid  %+v", 400, o.Payload)
}

func (o *DeletePolicyPathInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePolicyPathNotFound creates a DeletePolicyPathNotFound with default headers values
func NewDeletePolicyPathNotFound() *DeletePolicyPathNotFound {
	return &DeletePolicyPathNotFound{}
}

/*DeletePolicyPathNotFound handles this case with default header values.

Policy tree not found
*/
type DeletePolicyPathNotFound struct {
}

func (o *DeletePolicyPathNotFound) Error() string {
	return fmt.Sprintf("[DELETE /policy/{path}][%d] deletePolicyPathNotFound ", 404)
}

func (o *DeletePolicyPathNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePolicyPathFailure creates a DeletePolicyPathFailure with default headers values
func NewDeletePolicyPathFailure() *DeletePolicyPathFailure {
	return &DeletePolicyPathFailure{}
}

/*DeletePolicyPathFailure handles this case with default header values.

Error while deleting policy
*/
type DeletePolicyPathFailure struct {
	Payload models.Error
}

func (o *DeletePolicyPathFailure) Error() string {
	return fmt.Sprintf("[DELETE /policy/{path}][%d] deletePolicyPathFailure  %+v", 500, o.Payload)
}

func (o *DeletePolicyPathFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
