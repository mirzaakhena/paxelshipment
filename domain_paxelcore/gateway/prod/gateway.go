package prod

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"paxel-shipment/domain_paxelcore/model/entity"
	"paxel-shipment/domain_paxelcore/model/service"
	"paxel-shipment/shared/driver"
	"paxel-shipment/shared/infrastructure/config"
	"paxel-shipment/shared/infrastructure/logger"
	"time"
	// "github.com/ostafen/clover"
)

type gateway struct {
	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
	// db *clover.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, config *config.Config) *gateway {

	// db, err := clover.Open("database")
	// if err != nil {
	// 	panic(err.Error())
	// }
	//
	// exist, err := db.HasCollection("order")
	// if err != nil {
	// 	panic(err.Error())
	// }
	//
	// if !exist {
	// 	err = db.CreateCollection("order")
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }

	return &gateway{
		log:     log,
		appData: appData,
		config:  config,
		// db:      db,
	}
}

func (r *gateway) CallPaxelShipmentAPI(ctx context.Context, req *entity.Shipment, signature string) (*service.CallPaxelShipmentAPIServiceResponse, error) {
	r.log.Info(ctx, "called")

	byteJson, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := "https://stage-commerce-api.paxel.co/v1/shipments"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(byteJson))
	if err != nil {
		return nil, err
	}

	request.Header.Set("X-Paxel-API-Key", "ECOM.26A734G826MC034")
	request.Header.Set("X-Paxel-Signature", signature)
	request.Header.Set("Content-Type", "application/json")

	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	r.log.Info(ctx, "response from paxel %s", string(responseBody))

	return nil, nil
}

func (r *gateway) GetSecretKey(ctx context.Context) (string, error) {
	r.log.Info(ctx, "called")

	return "ECOM.8OQ3U04963GP231", nil
}
