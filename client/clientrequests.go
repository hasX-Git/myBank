package client

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	//"fmt"
	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	Firstn string `json:"add_cl_fn"`
	Lastn  string `json:"add_cl_ln"`
	NID    string `json:"add_cl_nid"`
}

func CreateAccount(c *gin.Context) {
	//initializing new client
	var newAccountRequest CreateAccountRequest
	if err := c.BindJSON(&newAccountRequest); err != nil {
		return
	}

	//checking if the client already exists
	for _, acc := range Clients {
		if newAccountRequest.NID == acc.NID {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "The account with such National ID already exists"})
			return
		}
	}

	//creating Client and Account
	var newClient ClientInfo
	newClient.Firstn = newAccountRequest.Firstn
	newClient.Lastn = newAccountRequest.Lastn
	newClient.NID = newAccountRequest.NID

	Clients = append(Clients, newClient)

	var newAccount Account
	newAccount.Personinfo = newClient
	newAccount.Balance = 0
	//generating unique Account ID(AID)
	var aid string
	aid += strconv.Itoa(time.Now().Year()) + strconv.Itoa(int(time.Now().Month())) + strconv.Itoa(time.Now().Day()) + strconv.Itoa(rand.Intn(999999-100000+1)+100000)
	newAccount.AID = aid

	Accounts = append(Accounts, newAccount)

	c.IndentedJSON(http.StatusCreated, newAccount)
}
