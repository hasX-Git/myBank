package main

import (
	"MyBankProject/client"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	client.ConnectToDB()
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{"message": "up"})
	})

	router.POST("/bank/account/create/", client.POSTcreateAccount)
	router.POST("/bank/account/pay/", client.POSTcreateTransaction)
	router.POST("/bank/upload", client.POSTfile)

	router.GET("/bank/find/account/:account/", client.GETaccountInfoByAID)
	router.GET("/bank/find/transaction/:transaction/", client.GETtransactionInfoByTID)
	router.GET("/bank/find/client/:id/", client.GETclientInfoByNID)
	router.GET("/bank/list/accounts/", client.GETaccountsList)
	router.GET("/bank/list/clients/", client.GETclientsList)
	router.GET("/bank/list/transactions/", client.GETtransactionsList)
	router.GET("/bank/download", client.GETexcelFile)
	router.GET("/bank/download/:file", client.GETfile)

	router.PATCH("/bank/account/deposit", client.PATCHdepositMoney)
	router.PATCH("/bank/deleteAll/", client.PATCHhardDeleteAll)

	port := ":" + os.Getenv("APP_PORT")
	router.Run(port) //AP
}
