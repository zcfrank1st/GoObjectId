package service

import (
    "github.com/gin-gonic/gin"
    "define"
    "fmt"
)

func GetObjectId(context *gin.Context) {
    objBytes := define.ObjectId()
    context.JSON(200, gin.H{
        "objectId": fmt.Sprintf("%x", objBytes),
        "objectIdString": string(objBytes),
        "objectIdBytes": objBytes,
    })
}