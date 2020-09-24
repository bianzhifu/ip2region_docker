package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

var (
	region *ip2region.Ip2Region
)

func index(c *gin.Context) {
	ip := c.ClientIP()
	ipRegion, _ := region.MemorySearch(ip)
	c.String(200, "当前IP:"+ip+",来自于:"+ipRegion.Country+"-"+ipRegion.Province+"-"+ipRegion.City+"-"+ipRegion.ISP)
}

func findByIp(c *gin.Context) {
	ip := c.Param("ip")
	ipRegion, _ := region.MemorySearch(ip)
	c.JSON(200, ipRegion)
}

func main() {
	//init ip2region
	var err error
	region, err = ip2region.New("/data/ip2region.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer region.Close()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/", index)
	r.GET("/ip=:ip", findByIp)
	r.Run(":8080")
}
