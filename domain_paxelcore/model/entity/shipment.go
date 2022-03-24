package entity

import (
	"crypto/sha256"
	"fmt"
)

type Shipment struct {
	InvoiceValue         int    `json:"invoice_value"`
	Origin               Node   `json:"origin"`
	Destination          Node   `json:"destination"`
	PaymentType          string `json:"payment_type"`
	InvoiceNumber        string `json:"invoice_number"`
	Items                []Item `json:"items"`
	IsCustomDeliveryTime bool   `json:"is_custom_delivery_time"`
	PickupDatetime       string `json:"pickup_datetime"`
	DeliveryDatetime     string `json:"delivery_datetime"`
	NeedInsurance        bool   `json:"need_insurance"`
	Note                 string `json:"note"`
}

type Node struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Note      string `json:"note"`
	Longitude int    `json:"longitude"`
	Latitude  int    `json:"latitude"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"district"`
	ZipCode   string `json:"zip_code"`
}

type Item struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	IsFragile bool   `json:"is_fragile"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Weight    int    `json:"weight"`
	Length    int    `json:"length"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type ShipmentRequest struct {
}

func NewShipment(req ShipmentRequest) (*Shipment, error) {

	var obj Shipment

	// assign value here

	err := obj.Validate()
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (r *Shipment) Validate() error {
	return nil
}

func (r Shipment) GetSignature(secretkey string) string {
	x := fmt.Sprintf("%s%s%s%s%s", r.InvoiceNumber[0:2], r.Origin.Name[0:2], r.Destination.Name[0:2], r.Items[0].Name[0:2], secretkey)

	sha256 := sha256.Sum256([]byte(x))

	return fmt.Sprintf("%x", sha256)

}
