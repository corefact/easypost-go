package easypost

/*
Official API documentation available at:
https://www.easypost.com/docs/api.html#orders
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

//Order is an Easypost order
type Order struct {
	ID        string     `json:"id,omitempty"`
	Object    string     `json:"object,omitempty"`
	Mode      string     `json:"mode,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	Reference     string      `json:"reference,omitempty"`
	ToAddress     *Address    `json:"to_address,omitempty"`
	FromAddress   *Address    `json:"from_address,omitempty"`
	ReturnAddress *Address    `json:"return_address,omitempty"`
	BuyerAddress  *Address    `json:"buyer_address,omitempty"`
	Shipments     []*Shipment `json:"shipments,omitempty"`
	Rates         Rates       `json:"rates,omitempty"`
	Messages      []*Message  `json:"messages,omitempty"`
	IsReturn      bool        `json:"is_return,omitempty"`
	Options       *Options    `json:"options,omitempty"`

	Carrier string `json:"carrier,omitempty"`
	Service string `json:"service,omitempty"`

	Error *Error `json:"error,omitempty"`

	retries int
}

//ErrorString returns a string representation of the error
func (o Order) ErrorString() string {
	if o.Error == nil {
		return ""
	}
	var epErr = fmt.Sprintf("%v (%v)", o.Error.Message, o.Error.Code)
	if o.Error.Code == ErrorOrderRateUnavailable {
		epErr = fmt.Sprintf("%v. Rates: ", epErr)
		if len(o.Rates) == 0 {
			epErr = fmt.Sprintf("%v None available ", epErr)
		} else {
			for _, r := range o.Rates {
				epErr = fmt.Sprintf("%v [%v] %v (%v %v).", epErr, r.Carrier, r.Service, r.Rate, r.Currency)
			}
		}
	}
	return epErr
}

//Create an EasyPost order
func (o *Order) Create() error {
	obj, err := Request.do("POST", "order", "", nest("order", o))

	if err != nil {
		return err
	}

	return json.Unmarshal(obj, &o)
}

//Buy an EasyPost order
func (o *Order) Buy() error {
	if o.Carrier == "" {
		return errors.New("carrier is missing")
	}
	if o.Service == "" {
		return errors.New("service rate is missing")
	}

	obj, err := Request.do("POST", "order", fmt.Sprintf("%v/buy", o.ID), map[string]string{
		"carrier": o.Carrier,
		"service": o.Service,
	})

	if err != nil {
		return err
	}

	if err := json.Unmarshal(obj, &o); err != nil {
		return err
	}

	if o.retries < 2 {
		o.retries++

		if o.Error != nil && o.Error.Code == ErrorOrderRateUnavailable {
			return o.buyBestRate()
		}
	}

	return nil
}

var services = map[string][][]string{
	"fedex": {
		[]string{"FEDEX_GROUND", "STANDARD_OVERNIGHT", "GROUND_HOME_DELIVERY"},
		[]string{"FEDEX_2_DAY_AM", "FEDEX_2_DAY", "FEDEX_EXPRESS_SAVER"},
		[]string{"FIRST_OVERNIGHT", "PRIORITY_OVERNIGHT"},
	},
	"ups": {
		[]string{"Ground"},
		[]string{"3rdDayAir", "UPSSaver", "UPS1DaySaver"},
		[]string{"NextDayAir", "STANDARD_OVERNIGHT", "NextDayAirEarlyAM", "PRIORITY_OVERNIGHT"},
	},
}

func serviceGroupIndex(service string, group [][]string) int {
	for i := 0; i < len(group); i++ {

		group := group[i]
		for _, s := range group {
			if s == service {
				return i
			}
		}
	}

	return 0
}

func (o *Order) buyBestRate() error {
	DebugLog("service not available %v", o.Service)
	o.Service = o.bestServiceMatch()
	if o.Service == "" {
		DebugLog("could not find an alternative service")
		return fmt.Errorf(o.ErrorString())
	}
	DebugLog("retrying buy with %v", o.Service)
	o.Error = nil
	return o.Buy()
}

func (o *Order) bestServiceMatch() string {

	availableServices := o.Rates.Services("UPS", "FedEx")
	serviceGroups := services[strings.ToLower(o.Carrier)]

	DebugLog("available services: %v", availableServices)
	isAvailable := stringSliceContains(availableServices)
	breakPoint := serviceGroupIndex(o.Service, serviceGroups)
	servicesOrderedByPreference := serviceGroups[breakPoint:]
	servicesOrderedByPreference = append(servicesOrderedByPreference, serviceGroups[:breakPoint]...)

	for _, services := range servicesOrderedByPreference {
		for _, service := range services {
			if service != o.Service && isAvailable(service) {
				return service
			}
		}
	}

	return ""
}

func stringSliceContains(s []string) func(string) bool {
	return func(t string) bool {
		for _, u := range s {
			if u == t {
				return true
			}
		}

		return false
	}
}
