package easypost

/*
Official API documentation available at:
https://www.easypost.com/docs/api.html#options
*/

//Options is an EasyPost object and defines the list of carrier-specific options for a shipment
type Options struct {
	AdditionalHandling       bool   `json:"additional_handling,omitempty"`
	AddressValidationLevel   string `json:"address_validation_level,omitempty"`
	Alcohol                  bool   `json:"alchol,omitempty"`
	BillReceiverAccount      string `json:"bill_receiver_account,omitempty"`
	BillReceiverPostalCode   string `json:"bill_receiver_postal_code,omitempty"`
	BillThirdPartyAccount    string `json:"bill_third_party_account,omitempty"`
	BillThirdPartyCountry    string `json:"bill_third_party_country,omitempty"`
	BillThirdPartyPostalCode string `json:"bill_third_party_postal_code,omitempty"`
	ByDrone                  bool   `json:"by_drone,omitempty"`
	CarbonNeutral            bool   `json:"carbon_neutral,omitempty"`
	CodAmount                string `json:"cod_amount,omitempty"`
	Currency                 string `json:"currency,omitempty"`
	DeliveredDutyPaid        bool   `json:"delivery_duty_paid,omitempty"`
	DeliveryConfirmation     string `json:"delivery_confirmation,omitempty"`
	DryIce                   bool   `json:"dry_ice,omitempty"`
	DryIceMedical            string `json:"dry_ice_medical,omitempty"`
	DryIceWeight             string `json:"dry_icd_weight,omitempty"`
	FreightCarge             int    `json:"freight_charge,omitempty"`
	HandlingInstructions     string `json:"handling_instructions,omitempty"`
	Hazmat                   string `json:"hazmat,omitempty"`
	HoldForPickup            bool   `json:"hold_for_pickup,omitempty"`
	InvoiceNumber            string `json:"invoice_number,omitempty"`
	LabelDate                string `json:"label_date,omitempty"`
	LabelFormat              string `json:"label_format,omitempty"`
	Machinable               bool   `json:"machinable,omitempty"`
	PrintCustom1             string `json:"print_custom_1,omitempty"`
	PrintCustom2             string `json:"print_custom_2,omitempty"`
	PrintCustom3             string `json:"print_custom_3,omitempty"`
	PrintCustom1Barcode      string `json:"print_custom_1_barcode,omitempty"`
	PrintCustom2Barcode      string `json:"print_custom_2_barcode,omitempty"`
	PrintCustom3Barcode      string `json:"print_custom_3_barcode,omitempty"`
	PrintCustom1Code         string `json:"print_custom_1_code,omitempty"`
	PrintCustom2Code         string `json:"print_custom_2_code,omitempty"`
	PrintCustom3Code         string `json:"print_custom_3_code,omitempty"`
	SaturdayDelivery         bool   `json:"saturday_delivery,omitempty"`
	SpecialRatesEligibility  string `json:"special_rates_eligibility,omitempty"`
	SmartpostHub             string `json:"smartpost_hub,omitempty"`
	SmartpostManifest        string `json:"smartpost_manifest,omitempty"`
}
