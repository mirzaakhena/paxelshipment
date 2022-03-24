package payload

import (
	"paxel-shipment/shared/driver"
)

type Payload struct {
	Data      interface{}            `json:"data"`
	Publisher driver.ApplicationData `json:"publisher"`
	TraceID   string                 `json:"traceId"`
}
