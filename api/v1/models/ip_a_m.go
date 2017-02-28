package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// IPAM IPAM configuration of an endpoint
// swagger:model IPAM
type IPAM struct {

	// endpoint
	// Required: true
	Endpoint *EndpointAddressing `json:"endpoint"`

	// host addressing
	// Required: true
	HostAddressing *NodeAddressing `json:"host-addressing"`
}

// Validate validates this IP a m
func (m *IPAM) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHostAddressing(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAM) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	if m.Endpoint != nil {

		if err := m.Endpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoint")
			}
			return err
		}
	}

	return nil
}

func (m *IPAM) validateHostAddressing(formats strfmt.Registry) error {

	if err := validate.Required("host-addressing", "body", m.HostAddressing); err != nil {
		return err
	}

	if m.HostAddressing != nil {

		if err := m.HostAddressing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host-addressing")
			}
			return err
		}
	}

	return nil
}