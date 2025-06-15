package main

import (
	"MyBankProject/client"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/bank/account", client.CreateAccount)
	router.POST("/bank/account/transaction", client.CreateTransaction)
	router.PATCH("bank/account/balance", client.DepositMoney)
	router.GET("/bank/:account", client.GetAccountInfoByAID)
	router.GET("bank/account/:transaction", client.GetTransactionInfoByTID)
	router.GET("bank/lists/accounts", client.GetAccountsList)
	router.GET("bank/lists/clients", client.GetClientsList)
	router.GET("bank/lists/transactions", client.GetTransacitonsList)
	router.Run("localhost:2266")
}
