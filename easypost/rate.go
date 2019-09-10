package easypost

/*
Official API documentation available at:
https://www.easypost.com/docs/api.html#rates
*/

import "time"

const (
	ErrorOrderRateUnavailable = "ORDER.RATE.UNAVAILABLE"
)

//Rate is an EasyPost object and defines the shipment rate, fetched after shipment creation
type Rate struct {
	ID        string     `json:"id,omitempty"`
	Object    string     `json:"object,omitempty"`
	Mode      string     `json:"mode,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	Service                string     `json:"service,omitempty"`
	Carrier                string     `json:"carrier,omitempty"`
	CarrierAccountID       string     `json:"carrier_account_id,omitempty"`
	ShipmentID             string     `json:"shipment_id,omitempty"`
	Rate                   string     `json:"rate,omitempty"`
	Currency               string     `json:"currency,omitempty"`
	RetailRate             string     `json:"retail_rate,omitempty"`
	RetailCurrency         string     `json:"retail_currency,omitempty"`
	ListRate               string     `json:"list_rate,omitempty"`
	ListCurrency           string     `json:"list_currency,omitempty"`
	DeliveryDays           int        `json:"delivery_days,omitempty"`
	DeliveryDate           *time.Time `json:"delivery_date,omitempty"`
	DaliveryDateGuaranteed bool       `json:"delivery_date_guaranteed,omitempty"`
}

type Rates []*Rate

func (r Rates) Filter(p func(*Rate) bool) Rates {
	f := Rates{}
	for _, rate := range r {
		if p(rate) {
			f = append(f, rate)
		}
	}

	return f
}

func (r Rates) Services(carrier string) []string {
	services := []string{}
	for _, rate := range r {
		if rate.Carrier == carrier {
			services = append(services, rate.Service)
		}
	}
	return services
}
