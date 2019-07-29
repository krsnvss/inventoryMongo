package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.POST("/pc/add", addPC)
	v1.POST("/pc/update", updatePC)
	v1.GET("/pc/all", getAllPC)
	v1.GET("/pc", getOnePC)
	v1.POST("/electronics/add", addElectronics)
	v1.POST("/electronics/update", updateElectronics)
	v1.POST("/network/add", addNetworkDevice)
	v1.POST("/network/update", updateNetworkDevice)
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", homePage)
	router.GET("/pc/all", allPC)
	router.POST("/device/add", addDevice)
	router.GET("/device/all", allDevices)
	router.GET("/device/network/all", allNetworkDevices)
	router.Run()
}
