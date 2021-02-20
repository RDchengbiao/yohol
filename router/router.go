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
	
	courseGroup := router.Group("/api/course", m.AddNotice("customerNotice", "v1"))
	{
		courseGroup.POST("/buy", course.BuyCourse)
		courseGroup.GET("/getinfo", course.GetCourseInfo)
	}

}

func login(c *gin.Context){



}

func loginv2(c *gin.Context){



}