package router

import (
	cardD "autfinal/internal/microservice/card/delivery"

	"github.com/gin-gonic/gin"
	"autfinal/internal/middleware"
)

func CardEndpoints(r *gin.RouterGroup, mws *middleware.Middlewares, cD *cardD.CardDelivery) {
	r.POST("", mws.MiddlewareCardFormData(), cD.CreateCard)
	r.GET("", cD.GetCards)
}