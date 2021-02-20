package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine)  {

	// Simple group: v1
	v1 := r.Group("/v1")
	{
		v1.POST("/login", login)
	}

	// Simple group: v2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", loginv2)
	}
}

func login(c *gin.Context){



}

func loginv2(c *gin.Context){



}