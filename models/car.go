//
// Copyright Uhuchain All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0
//

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

// Car car
// swagger:model Car

type Car struct {

	// brand
	// Required: true
	Brand *string `json:"brand"`

	// id
	ID int64 `json:"id,omitempty"`

	// model
	// Required: true
	Model *string `json:"model"`

	// policies
	Policies []*Policy `json:"policies"`

	// vehicle Id
	// Required: true
	VehicleID *string `json:"vehicleId"`
}

/* polymorph Car brand false */

/* polymorph Car id false */

/* polymorph Car model false */

/* polymorph Car policies false */

/* polymorph Car vehicleId false */

// Validate validates this car
func (m *Car) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrand(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVehicleID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Car) validateBrand(formats strfmt.Registry) error {

	if err := validate.Required("brand", "body", m.Brand); err != nil {
		return err
	}

	return nil
}

func (m *Car) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

func (m *Car) validatePolicies(formats strfmt.Registry) error {

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

func (m *Car) validateVehicleID(formats strfmt.Registry) error {

	if err := validate.Required("vehicleId", "body", m.VehicleID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Car) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Car) UnmarshalBinary(b []byte) error {
	var res Car
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
