package easypost

/*
Official API documentation available at:
https://www.easypost.com/docs/api.html#parcels
*/

import (
	"time"
)

//Parcel is an EasyPost object that defines a shipping parcel
type Parcel struct {
	ID        string     `json:"id,omitempty"`
	Object    string     `json:"object,omitempty"`
	Mode      string     `json:"mode,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	Length            float32 `json:"length,omitempty"`
	Width             float32 `json:"width,omitempty"`
	Height            float32 `json:"height,omitempty"`
	PredefinedPackage string  `json:"predefined_package,omitempty"`
	Weight            float32 `json:"weight,omitempty"`
}
