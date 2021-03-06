// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListenerTemplate listener template
// swagger:model ListenerTemplate
type ListenerTemplate struct {

	// certificate instance
	CertificateInstance *ListenerTemplateCertificateInstance `json:"certificate_instance,omitempty"`

	// The connection limit of the listener.
	// Maximum: 15000
	// Minimum: 1
	ConnectionLimit int64 `json:"connection_limit,omitempty"`

	// default pool
	DefaultPool *ListenerTemplateDefaultPool `json:"default_pool,omitempty"`

	// The list of policies of this listener
	Policies []*ListenerPolicyTemplate `json:"policies"`

	// The listener port number.
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Port *int64 `json:"port"`

	// The listener protocol.
	// Required: true
	// Enum: [http https tcp]
	Protocol *string `json:"protocol"`
}

// Validate validates this listener template
func (m *ListenerTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificateInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultPool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListenerTemplate) validateCertificateInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.CertificateInstance) { // not required
		return nil
	}

	if m.CertificateInstance != nil {
		if err := m.CertificateInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate_instance")
			}
			return err
		}
	}

	return nil
}

func (m *ListenerTemplate) validateConnectionLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionLimit) { // not required
		return nil
	}

	if err := validate.MinimumInt("connection_limit", "body", int64(m.ConnectionLimit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("connection_limit", "body", int64(m.ConnectionLimit), 15000, false); err != nil {
		return err
	}

	return nil
}

func (m *ListenerTemplate) validateDefaultPool(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultPool) { // not required
		return nil
	}

	if m.DefaultPool != nil {
		if err := m.DefaultPool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_pool")
			}
			return err
		}
	}

	return nil
}

func (m *ListenerTemplate) validatePolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	for i := 0; i < len(m.Policies); i++ {
		if swag.IsZero(m.Policies[i]) { // not required
			continue
		}

		if m.Policies[i] != nil {
			if err := m.Policies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListenerTemplate) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

var listenerTemplateTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","https","tcp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listenerTemplateTypeProtocolPropEnum = append(listenerTemplateTypeProtocolPropEnum, v)
	}
}

const (

	// ListenerTemplateProtocolHTTP captures enum value "http"
	ListenerTemplateProtocolHTTP string = "http"

	// ListenerTemplateProtocolHTTPS captures enum value "https"
	ListenerTemplateProtocolHTTPS string = "https"

	// ListenerTemplateProtocolTCP captures enum value "tcp"
	ListenerTemplateProtocolTCP string = "tcp"
)

// prop value enum
func (m *ListenerTemplate) validateProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, listenerTemplateTypeProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ListenerTemplate) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListenerTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListenerTemplate) UnmarshalBinary(b []byte) error {
	var res ListenerTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
