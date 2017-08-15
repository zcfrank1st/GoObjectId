package main

import (
    "github.com/gin-gonic/gin"
    "service"
)

func main() {
    serv := gin.Default()

    serv.GET("/status", func(c *gin.Context) {
        c.String(200, "OK")
    })

    g := serv.Group("/object-id")
    {
        g.GET("", service.GetObjectId)
    }

    serv.Run()
}
