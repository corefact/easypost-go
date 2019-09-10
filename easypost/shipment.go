package easypost

/*
Official API documentation available at:
https://www.easypost.com/docs/api.html#shipments
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

//these are all the possible label formats
const (
	LabelFormatEPL2 = "EPL2"
	LabelFormatPDF  = "PDF"
	LabelFormatPNG  = "PNG"
	LabelFormatZPL  = "ZPL"
)

//Shipment is an EasyPost object that defines a shipment
type Shipment struct {
	ID        string     `json:"id,omitempty"`
	Object    string     `json:"object,omitempty"`
	Mode      string     `json:"mode,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	Reference     string        `json:"reference,omitempty"`
	ToAddress     *Address      `json:"to_address,omitempty"`
	FromAddress   *Address      `json:"from_address,omitempty"`
	ReturnAddress *Address      `json:"return_address,omitempty"`
	BuyerAddress  *Address      `json:"buyer_address,omitempty"`
	Parcel        *Parcel       `json:"parcel,omitempty"`
	Rates         Rates         `json:"rates,omitempty"`
	SelectedRate  *Rate         `json:"selected_rate,omitempty"`
	PostageLabel  *PostageLabel `json:"postage_label,omitempty"`
	Messages      []*Message    `json:"messages,omitempty"`
	Options       *Options      `json:"options,omitempty"`
	IsReturn      bool          `json:"is_return,omitempty"`
	TrackingCode  string        `json:"tracking_code,omitempty"`
	UspsZone      int           `json:"usps_zone,omitempty"`
	Status        string        `json:"status,omitempty"`

	Error *Error `json:"error,omitempty"`
}

//Message is an EasyPost object that defines the message for a shipment
type Message struct {
	Carrier string `json:"carrier,omitempty"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

type PostageLabel struct {
	ID        string     `json:"id,omitempty"`
	Object    string     `json:"object,omitempty"`
	Mode      string     `json:"mode,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	IntegratedForm  string     `json:"integrated_form,omitempty"`
	LabelDate       *time.Time `json:"label_date,omitempty"`
	LabelEp12Url    string     `json:"label_ep12_url,omitempty"`
	LabelFileType   string     `json:"label_file_type,omitempty"`
	LabelPdfURL     string     `json:"label_pdf_url,omitempty"`
	LabelResolution int        `json:"label_resolution,omitempty"`
	LabelSize       string     `json:"label_size,omitempty"`
	LabelType       string     `json:"label_type,omitempty"`
	LabelURL        string     `json:"label_url,omitempty"`
	LabelZplURL     string     `json:"label_zpl_url,omitempty"`
}

//Create an EasyPost shipment
func (s *Shipment) Create() error {
	obj, err := Request.do("POST", "shipment", "", map[string]interface{}{"shipment": s})
	if err != nil {
		return errors.New("failed to request EasyPost shipment creation")
	}
	return json.Unmarshal(obj, &s)
}

//Buy an EasyPost shipment
func (s *Shipment) Buy() error {
	rate := s.SelectedRate.Service
	carrier := s.SelectedRate.Carrier
	if rate == "" {
		return errors.New("no rate has been selected")
	}
	if s.ID == "" {
		if err := s.Create(); err != nil {
			return err
		}
	}

	if s.SelectedRate.ID == "" {

		for _, singleRate := range s.Rates {
			if rate == singleRate.Service && carrier == singleRate.Carrier {
				s.SelectedRate = singleRate
				break
			}
		}
		if s.SelectedRate.ID == "" {
			return errors.New("Cannot find rate '" + rate + "' for the given carrier '" + carrier + "'")
		}

	}

	obj, err := Request.do("POST", "shipment", fmt.Sprintf("%v/buy", s.ID), fmt.Sprintf("rate[id]=%v", s.SelectedRate.ID))
	if err != nil {
		return errors.New("failed to request EasyPost shipment purchase")
	}

	err = json.Unmarshal(obj, &s)
	if err != nil {
		return errors.New("failed to decode EasyPost shipment purchase")
	}

	if s.Error != nil {
		return fmt.Errorf("failed to request EasyPost shipment purcahse: %v", s.Error.Message)
	}

	return nil
}
