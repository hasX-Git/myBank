package client

import (
	"net/http"
	//"time"
	//"fmt"
	"errors"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Transaction uint32 `json:"transaction"`
}

type Client struct {
	Firstname    string        `json:"firstname"`
	Lastname     string        `json:"lastname"`
	ClientID     string        `json:"clientid"`
	Balance      int32         `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}

var Clients []Client

func FindClientByID(id string) (*Client, error) {
	for index, cl := range Clients {
		if id == cl.ClientID {
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

// func AddTransaction(c *gin.Context, id string) {
// 	var newT Transaction

// 	if client, err := FindClientByID(id); err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The person not found"})
// 	}

// }
