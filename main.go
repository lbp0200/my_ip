package main

import (
	"github.com/gin-gonic/gin"
	geoIp2 "github.com/oschwald/geoip2-golang"
	"log"
	"os"
)

type GeoClient struct {
	IpInfoReader *geoIp2.Reader
}

func main() {
	InitGeoIp2Client()
	r := gin.Default()
	r.GET("/ip/:ip", func(c *gin.Context) {
		ipStr := c.Param("ip")
		if ipStr == "" {
			ipStr = c.ClientIP()
		}
		ipInfo, err := geoClient.GeoIp2GetIpInfo(ipStr)
		if err != nil {
			c.JSON(200, gin.H{"error": err})
			return
		}
		c.JSON(200, ipInfo)
	})
	r.GET("/my_ip", func(c *gin.Context) {
		ipInfo, err := geoClient.GeoIp2GetIpInfo(c.ClientIP())
		if err != nil {
			c.JSON(200, gin.H{"error": err})
			return
		}
		c.JSON(200, ipInfo)
	})
	log.Fatal(r.Run(os.Getenv("MY_LISTEN")))
}
