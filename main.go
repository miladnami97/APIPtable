package main

import (
	"log"
	"main/api"
	"main/util"

	"github.com/gin-gonic/gin"

	_ "main/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	hasPermission, err := util.CheckIptablesPermission()
	if err != nil {
		log.Fatalln("ERROR -", err)
	}

	if hasPermission == false {
		log.Fatalln("ERROR -", err)
	}

	router := gin.Default()

	router.POST("/firewall/rules", api.AddFirewallRule)
	router.DELETE("/firewall/rules", api.DeleteFirewallRules)

	router.POST("/firewall/ipset", api.CreateIPSetGroup)
	router.DELETE("/firewall/ipset", api.DeleteIPSetGroup)

	router.POST("/firewall/ipset/ip", api.AddIPToIPSet)
	router.DELETE("/firewall/ipset/ip", api.RemoveIPFromIPSet)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8089")
}
