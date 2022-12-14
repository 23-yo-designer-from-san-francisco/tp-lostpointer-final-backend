package middleware

import (
	"github.com/gin-gonic/gin"
	"autfinal/internal/models"
	"github.com/goccy/go-json"
	"log"
)

var allowedOrigins = []string{"", "http://45.141.102.243:8080", "http://127.0.0.1:8080", 
	"http://localhost:5173", "http://localhost:5174", "http://127.0.0.1:5173", "http://127.0.0.1:5174"}

type Middlewares struct {
	
}

func NewMiddleware() *Middlewares {
	return &Middlewares{

	}
}

func (m *Middlewares) MiddlewareCardFormData() gin.HandlerFunc {
	return func(c *gin.Context) {
		inputCard := c.Request.FormValue("json")
		card := new(models.CardDay)
		json.Unmarshal([]byte(inputCard), &card)

		c.Set("card", *card)
		c.Next()
	}
}

func (m *Middlewares) CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		isAllowed := false
		for _, orig := range allowedOrigins {
			if origin == orig {
				isAllowed = true
			}
		}

		if !isAllowed {
			log.Print("CORS not allowed origin = ", origin)
			return
		}

        c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}