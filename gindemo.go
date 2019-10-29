package main

import (
    "github.com/gin-gonic/gin"
    //"log"
    //"net/http"
)

func main(){
    //gin router with default middleware
    //BY default it serves on :8080
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context){
        c.JSON(200, gin.H{
             "message" : "pong",
           })
})
    r.Run()
    //router.Run(":3000") for a hard coded port
}
