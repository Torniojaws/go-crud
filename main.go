package main

import (
    "github.com/gin-gonic/gin"
    "go-crud/apps/hello"
)

func main() {
    router := gin.Default()
    
    v1 := router.Group("v1")
    {
        v1.POST("/hello", hello.Create)
        v1.GET("/hello", hello.Read)
        v1.PUT("/hello", hello.Update)
        v1.DELETE("/hello", hello.Delete)
    }
    
    router.Run(":8080")
}
