package createshipment

import "paxel-shipment/domain_paxelcore/model/service"

// Outport of usecase
type Outport interface {
	service.CallPaxelShipmentAPIService
	service.GetSecretKeyService
}
