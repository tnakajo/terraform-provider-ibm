// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostInstancesParamsBodyNetworkInterfacesItems NetworkInterfaceTemplate
//
// network interface
// swagger:model postInstancesParamsBodyNetworkInterfacesItems
type PostInstancesParamsBodyNetworkInterfacesItems struct {

	// The user-defined name for this network interface
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// The primary IPv4 address
	PrimaryIPV4Address string `json:"primary_ipv4_address,omitempty"`

	// Collection seconary IP addresses
	SecondaryAddresses []string `json:"secondary_addresses,omitempty"`

	// security groups
	SecurityGroups []*PostInstancesParamsBodyNetworkInterfacesItemsSecurityGroupsItems `json:"security_groups,omitempty"`

	// subnet
	Subnet *PostInstancesParamsBodyNetworkInterfacesItemsSubnet `json:"subnet,omitempty"`
}

// Validate validates this post instances params body network interfaces items
func (m *PostInstancesParamsBodyNetworkInterfacesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostInstancesParamsBodyNetworkInterfacesItems) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PostInstancesParamsBodyNetworkInterfacesItems) validateSecurityGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityGroups); i++ {
		if swag.IsZero(m.SecurityGroups[i]) { // not required
			continue
		}

		if m.SecurityGroups[i] != nil {
			if err := m.SecurityGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("security_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostInstancesParamsBodyNetworkInterfacesItems) validateSubnet(formats strfmt.Registry) error {

	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if m.Subnet != nil {
		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostInstancesParamsBodyNetworkInterfacesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostInstancesParamsBodyNetworkInterfacesItems) UnmarshalBinary(b []byte) error {
	var res PostInstancesParamsBodyNetworkInterfacesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
