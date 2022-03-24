package restapi

import (
	"github.com/gin-gonic/gin"

	"paxel-shipment/domain_paxelcore/usecase/createshipment"
	"paxel-shipment/shared/infrastructure/config"
	"paxel-shipment/shared/infrastructure/logger"
)

type Controller struct {
	Router               gin.IRouter
	Config               *config.Config
	Log                  logger.Logger
	CreateShipmentInport createshipment.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.GET("/createshipment", r.authorized(), r.createShipmentHandler(r.CreateShipmentInport))
}
