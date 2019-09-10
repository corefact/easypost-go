package easypost

/**
Official API documentation available at:
https://www.easypost.com/docs/api.html#addresses
**/

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	DefaultNullPhone = "0000000000"
)

//Address is ah EasyPost object that defines a shipping address
type Address struct {
	ID        string     `json:"id,omitempty"`
	Object    string     `json:"object,omitempty"`
	Mode      string     `json:"mode,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	Street1         string         `json:"street1,omitempty"`
	Street2         string         `json:"street2,omitempty"`
	City            string         `json:"city,omitempty"`
	State           string         `json:"state,omitempty"`
	Zip             string         `json:"zip,omitempty"`
	Country         string         `json:"country,omitempty"`
	Residential     bool           `json:"residential,omitempty"`
	CarrierFacility string         `json:"carrier_facility,omitempty"`
	Name            string         `json:"name,omitempty"`
	Company         string         `json:"company,omitempty"`
	Phone           string         `json:"phone,omitempty"`
	Email           string         `json:"email,omitempty"`
	FederalTaxID    string         `json:"federal_tax_id,omitempty"`
	StateTaxID      string         `json:"state_tax_id,omitempty"`
	Verifications   *Verifications `json:"verifications,omitempty"`
	Verification    bool           `json:"verify,omitempty"`
}

type Verifications struct {
	Zip4     *Verification `json:"zip4,omitempty"`
	Delivery *Verification `json:"delivery,omitempty"`
}

type Verification struct {
	Success bool          `json:"success,omitempty"`
	Errors  []*FieldError `json:"errors,omitempty"`
}

//Create a new EasyPost address
func (a *Address) Create() error {
	obj, err := Request.do("POST", "address", "", map[string]interface{}{"address": a})
	if err != nil {
		return errors.New("failed to request EasyPost Address creation")
	}
	return json.Unmarshal(obj, &a)
}

//Get Retrieves the address from EasyPost
func (a *Address) Get() error {
	obj, _ := Request.do("GET", "address", a.ID, "")
	return json.Unmarshal(obj, &a)
}

//Verify creates and verifies an address in EasyPost
func (a *Address) Verify() error {
	obj, err := Request.do("POST", "address", "", map[string]interface{}{
		"address": a,
		"verify":  []string{"delivery"},
	})
	DebugLog("easypost address verification results: %v", string(obj))
	if err != nil {
		return errors.New("failed to request EasyPost Address creation")
	}

	err = json.Unmarshal(obj, a)

	if err == nil {
		if len(a.Verifications.Delivery.Errors) > 0 {
			e := []string{}
			for _, err := range a.Verifications.Delivery.Errors {
				e = append(e, fmt.Sprintf("%v: %v", err.Field, err.Message))
			}

			m := strings.Join(e, ";")
			DebugLog(m)
			err = errors.New(m)
		}
		if a.ID == "" {
			err = errors.New("couldn't retrieve an EasyPost ID")
		}
	}
	return err
}
