package client

import (
	"net/http"
	//"time"
	//"fmt"
	"errors"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Tr   uint32 `json:"transaction"`
	TrID string `json:"transactionid"`
}

type Account struct {
	Balance int32         `json:"balance"`
	Trs     []Transaction `json:"transactions"`
	AID     string        `json:"aid"`
}

type Client struct {
	Firstn string  `json:"firstname"`
	Lastn  string  `json:"lastname"`
	NID    string  `json:"nid"`
	Acc    Account `json:"account"`
}

var Clients []Client

func FindClientByID(id string) (*Client, error) {
	for index, cl := range Clients {
		if id == cl.NID {
			return &Clients[index], nil
		}
	}
	return nil, errors.New("Not Found")
}

func AddClient(c *gin.Context) {
	var newClient Client
	if err := c.BindJSON(&newClient); err != nil {
		return
	}
	Clients = append(Clients, newClient)
	c.IndentedJSON(http.StatusCreated, newClient)
}

func AddTransaction(c *gin.Context) {
	var newT Transaction

	if err := c.BindJSON(&newT); err != nil {
		return
	}

	if client, err := FindClientByID(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The person not found"})
	}

}
