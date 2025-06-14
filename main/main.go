package main

import (
	"MyBankProject/client"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/bank/account", client.CreateAccount)
	router.POST("/bank/account/transaction", client.CreateTransaction)
	router.GET("/bank/:account", client.GetAccountInfoByAID)
	router.GET("bank/account/:transaction", client.GetAccountInfoByAID)
	router.Run("localhost:2266")
}
