package service

import (
	"context"
	"paxel-shipment/domain_paxelcore/model/entity"
)

type CallPaxelShipmentAPIService interface {
	CallPaxelShipmentAPI(ctx context.Context, req *entity.Shipment, signature string) (*CallPaxelShipmentAPIServiceResponse, error)
}

type CallPaxelShipmentAPIServiceRequest struct {
}

type CallPaxelShipmentAPIServiceResponse struct {
}
type GetSecretKeyService interface {
	GetSecretKey(ctx context.Context) (string, error)
}

type GetSecretKeyServiceRequest struct {
}

type GetSecretKeyServiceResponse struct {
}
