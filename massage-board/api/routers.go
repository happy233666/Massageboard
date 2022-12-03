package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	v := r.Group("/users")
	{
		v.POST("/register", Register)
		v.POST("/login", Login)
		v.PUT("/ChangePassword")
	}
	r.Run()
}
