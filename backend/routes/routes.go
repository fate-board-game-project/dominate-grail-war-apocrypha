package routes

import (
    "github.com/gin-gonic/gin"
    "backend/internal/api/v1"
    "backend/middleware"
)

func SetupRoutes(r *gin.Engine) {
    v1Group := r.Group("/api/v1")
    {
        v1Group.POST("/register", v1.Register)
        v1Group.POST("/login", v1.Login)
        
        authGroup := v1Group.Group("/")
        authGroup.Use(middleware.JWT())
        // {
        //     authGroup.POST("/rooms", v1.CreateRoom)
        //     authGroup.GET("/rooms/:id", v1.GetRoom)
        //     authGroup.DELETE("/rooms/:id", v1.DeleteRoom)
        //     authGroup.POST("/rooms/:id/join", v1.JoinRoom)
        // }
    }
}
