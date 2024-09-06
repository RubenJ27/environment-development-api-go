package routes

import (
	"development-environment-api-go-manager/src/api/rest"
	"development-environment-api-go-manager/src/api/rest/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

// Asegúrate de importar la interfaz UserRepository desde el paquete rest
type UserRepository = rest.UserRepository

func GetServer(userRepo UserRepository) *gin.Engine {
	r := gin.Default()

	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Usar el middleware de timeout
	r.Use(middleware.TimeoutMiddleware())

	users := r.Group("/users")
	userHandlers := rest.NewUserHandlers(userRepo)
	{
		users.GET("/:id", userHandlers.ReadUser)
	}

	return r
}