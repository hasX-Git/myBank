package client

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func randWithRange(digits int) int {
	//digits is how much digits u want in a random number
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var min int = 1
	var max int = 9
	var i int
	for i = 0; i < digits-1; i++ {
		min = min * 10
		max = max*10 + 9
	}
	return r.Intn(max-min+1) + min
}

func currentDateAsID(n int) string {
	return strconv.Itoa(time.Now().Year() + int(time.Now().Month()) + time.Now().Day() + randWithRange(n))
}

func checkValidityOfNID(id string, n int) bool {
	if len(id) != n {
		return false
	}
	for _, ch := range id {
		if int(ch) < 48 && int(ch) > 57 {
			return false
		}
	}
	return true
}

type CreateAccountRequest struct {
	Firstn string `json:"add_cl_fn"`
	Lastn  string `json:"add_cl_ln"`
	NID    string `json:"add_cl_nid"`
}

func CreateAccount(c *gin.Context) {
	//initializing new client
	var newAccountRequest CreateAccountRequest
	if err := c.BindJSON(&newAccountRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message ": "Improper formatting"})
		return
	}

	//checking validity of NID
	if !checkValidityOfNID(newAccountRequest.NID, 12) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message ": "Invalid NID"})
		return
	}

	//checking if the client already exists
	_, isPresent := Clients[newAccountRequest.NID]
	if isPresent {
		c.IndentedJSON(http.StatusConflict, gin.H{"message ": "The account with such National ID already exists"})
		return
	}

	//creating Client and Account
	var newClient ClientInfo
	newClient.Firstn = newAccountRequest.Firstn
	newClient.Lastn = newAccountRequest.Lastn
	newClient.NID = newAccountRequest.NID
	Clients[newClient.NID] = &newClient

	var newAccount Account
	newAccount.Personinfo = newClient
	newAccount.Balance = 0
	newAccount.Trs = make(map[string]*Transaction)

	//generating unique Account ID(AID)
	var aid string
	for {
		aid = currentDateAsID(5)
		if _, isIn := Accounts[aid]; !isIn {
			break
		}
	}
	newAccount.AID = aid
	Accounts[newAccount.AID] = &newAccount

	//
	c.IndentedJSON(http.StatusCreated, newAccount)
}

type CreateTransactionRequest struct {
	Aid string `json:"add_tr_aid"`
	Sum uint32 `json:"add_tr_sum"`
}

func CreateTransaction(c *gin.Context) {
	var newTransactionRequest CreateTransactionRequest

	if err := c.BindJSON(&newTransactionRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Improper formatting"})
		return
	}

	_, isPresent := Accounts[newTransactionRequest.Aid]
	if !isPresent {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The account with such AID doesn't exist"})
		return
	}

	var newtr Transaction
	newtr.Sum = newTransactionRequest.Sum

	//checking if enough balance on account
	if newtr.Sum > Accounts[newTransactionRequest.Aid].Balance {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "not enough balance"})
		return
	}

	newtr.TrID = currentDateAsID(5)

	Transactions[newtr.TrID] = &newtr

	Accounts[newTransactionRequest.Aid].Trs[newtr.TrID] = &newtr
	Accounts[newTransactionRequest.Aid].Balance -= newtr.Sum
}
