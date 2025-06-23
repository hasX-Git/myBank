package main

import (
	"MyBankProject/client"
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	initializers.LoadEnvVars()
// 	initializers.DBconnection()
// }

func main() {

	connStr := "postgres://postgres:sctpwd@localhost:5432/mybank?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

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
