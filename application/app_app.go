package application

import (
	"paxel-shipment/domain_paxelcore/controller/restapi"
	"paxel-shipment/domain_paxelcore/gateway/prod"
	"paxel-shipment/domain_paxelcore/usecase/createshipment"
	"paxel-shipment/shared/driver"
	"paxel-shipment/shared/infrastructure/config"
	"paxel-shipment/shared/infrastructure/logger"
	"paxel-shipment/shared/infrastructure/server"
	"paxel-shipment/shared/infrastructure/util"
)

type app struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c app) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewApp() func() driver.RegistryContract {
	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData("app", appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandlerDefault(log, appData, cfg)

		datasource := prod.NewGateway(log, appData, cfg)

		return &app{
			httpHandler: &httpHandler,
			controller: &restapi.Controller{
				Log:                  log,
				Config:               cfg,
				Router:               httpHandler.Router,
				CreateShipmentInport: createshipment.NewUsecase(datasource),
			},
		}

	}
}
