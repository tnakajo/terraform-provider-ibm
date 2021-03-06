// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ListenerTemplateCertificateInstance The certificate instance used for SSL termination. It is applicable only to `https` protocol.
// swagger:model listenerTemplateCertificateInstance
type ListenerTemplateCertificateInstance struct {

	// The ceritificate instance's CRN
	Crn string `json:"crn,omitempty"`
}

// Validate validates this listener template certificate instance
func (m *ListenerTemplateCertificateInstance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListenerTemplateCertificateInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListenerTemplateCertificateInstance) UnmarshalBinary(b []byte) error {
	var res ListenerTemplateCertificateInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
