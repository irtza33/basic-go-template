package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	Server 	string
	DB 		string
}

func Health() gin.HandlerFunc { // TODO: DB health should ping DB
	return func(ctx *gin.Context) {
		healthCheck := HealthCheck{
			Server: "UP",
			DB: "UP",
		}

		ctx.JSON(http.StatusOK, healthCheck)
	}
}