package middleware

import (
	"github.com/gin-gonic/gin"
	"autfinal/internal/models"
	"github.com/goccy/go-json"
)

type Middlewares struct {
	
}

func NewMiddleware() *Middlewares {
	return &Middlewares{

	}
}

func (m *Middlewares) MiddlewareCardFormData() gin.HandlerFunc {
	return func(c *gin.Context) {
		inputCard := c.Request.FormValue("json")
		card := new(models.Card)
		json.Unmarshal([]byte(inputCard), &card)

		c.Set("card", *card)
		c.Next()
	}
}