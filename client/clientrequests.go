package client

import (
	"net/http"
	//"time"
	//"fmt"
	"errors"

	"github.com/gin-gonic/gin"
)

var AIDs uint32 = 0

type AddClientRequest struct {
	Firstn string `json:"add_cl_fn"`
	Lastn  string `json:"add_cl_ln"`
	NID    string `json:"add_cl_nid"`
}

func AddClient(c *gin.Context) {

	var newClientRequest AddClientRequest
	if err := c.BindJSON(&newClientRequest); err != nil {
		return
	}

	var newClient Client
	newClient.Firstn = newClientRequest.Firstn
	newClient.Lastn = newClientRequest.Lastn
	newClient.NID = newClientRequest.NID
	newClient.Acc.Balance = 0
	newClient.Acc.AID = string(AIDs)
	AIDs++

	Clients = append(Clients, newClient)
	c.IndentedJSON(http.StatusCreated, newClient)
}

func FindClientByID(id string) (*Client, error) {
	for index, cl := range Clients {
		if id == cl.NID {
			return &Clients[index], nil
		}
	}
	return nil, errors.New("Not Found")
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
