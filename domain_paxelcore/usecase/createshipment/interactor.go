package createshipment

import (
	"context"
	"paxel-shipment/domain_paxelcore/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type createShipmentInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &createShipmentInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *createShipmentInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	res := &InportResponse{}

	obj := entity.Shipment{
		InvoiceValue: 899000,
		Origin: entity.Node{
			Name:      "Yone",
			Email:     "",
			Phone:     "6281122334455",
			Address:   "Pasar Pagi Mangga Dua, Lt. 5 Blok B No. 53A - 73, Jl. Mangga Dua Raya, RT.11/RW.5, Ancol, Kec. Pademangan, Kota Jkt Utara, Daerah Khusus Ibukota Jakarta 14430, Indonesia",
			Note:      "",
			Longitude: 0,
			Latitude:  0,
			Province:  "",
			City:      "",
			District:  "",
			ZipCode:   "",
		},
		Destination: entity.Node{
			Name:      "Katarina",
			Email:     "katarina@paxel.co",
			Phone:     "6289932123421",
			Address:   "Jalan Cibaduyut Lama No.23A Kebon Lega, Kec. Bojongloa Kidul Bandung City, West Java, Indonesia 40235",
			Note:      "",
			Longitude: 0,
			Latitude:  0,
			Province:  "",
			City:      "",
			District:  "",
			ZipCode:   "",
		},
		PaymentType:   "CRD",
		InvoiceNumber: req.InvoiceNumber,
		Items: []entity.Item{
			{
				Code:      "SKU7823123",
				Name:      "Sepatu Nike",
				Category:  "Hobbies",
				IsFragile: false,
				Price:     899000,
				Quantity:  1,
				Weight:    2000,
				Length:    20,
				Width:     10,
				Height:    10,
			},
		},
		IsCustomDeliveryTime: false,
		PickupDatetime:       "2022-03-26 10:00:00",
		DeliveryDatetime:     "",
		NeedInsurance:        true,
		Note:                 "",
	}

	secretKey, err := r.outport.GetSecretKey(ctx)
	if err != nil {
		return nil, err
	}

	signature := obj.GetSignature(secretKey)

	_, err = r.outport.CallPaxelShipmentAPI(ctx, &obj, signature)
	if err != nil {
		return nil, err
	}

	return res, nil
}
