package easypost

import (
	"strings"

	"github.com/corefact/easypost-go/fake"
)

type RequestControllerFake struct {
	CreateOrderResponse func() []byte
	BuyOrderResponse    func() []byte
	AddressResponse     func() []byte
}

func (rif RequestControllerFake) do(method string, objectType string, objectUrl string, payload interface{}) ([]byte, error) {
	//Address
	if objectType == "address" {
		if rif.AddressResponse != nil {
			return rif.AddressResponse(), nil
		}
		//Address.Create()
		if objectUrl == "" {
			return []byte(fake.EasypostFakeAddress), nil
		}
		//Address.Get()
		if objectUrl != "" {
			return []byte(fake.EasypostFakeAddress), nil
		}
	}

	//Shipment
	if objectType == "shipment" {
		//Shipment.Create()
		if objectUrl == "" {
			return []byte(fake.ShipmentCreate), nil
		}
		//Shipment.Buy()
		if objectUrl != "" {
			return []byte(fake.ShipmentBuy), nil
		}
	}

	if objectType == "order" {
		if rif.CreateOrderResponse != nil && objectUrl == "" {
			return rif.CreateOrderResponse(), nil
		}

		if rif.BuyOrderResponse != nil && strings.Contains(objectUrl, "buy") {
			return rif.BuyOrderResponse(), nil
		}
		//Order.Create()
		if objectUrl == "" {
			return []byte(fake.EasypostFakeCreateOrder), nil
		}
		//Order.Buy()
		if objectUrl != "" {
			return []byte(fake.EasypostFakeBuyOrder), nil
		}
	}

	if objectType == "customs_info" {
		if objectUrl == "" {
			return []byte(fake.EasypostFakeCustoms), nil
		}
	}
	return nil, nil
}
