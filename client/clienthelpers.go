package client

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
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
	return strconv.Itoa(time.Now().Year()) + strconv.Itoa(int(time.Now().Month())) + strconv.Itoa(time.Now().Day()) + strconv.Itoa(randWithRange(n))
}

func checkValidityOfID(id string, n int) bool {
	if len(id) != n {
		return false
	}

	//only digits
	for _, ch := range id {
		if int(ch) < 48 && int(ch) > 57 {
			return false
		}
	}
	return true
}

func hash(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

type createAccountRequest struct {
	Firstn string `json:"add_cl_fn"`
	Lastn  string `json:"add_cl_ln"`
	NID    string `json:"add_cl_nid"`
}

type createTransactionRequest struct {
	AID string `json:"add_tr_aid"`
	Sum uint32 `json:"add_tr_sum"`
}

type depositRequest struct {
	AID string `json:"dep_aid"`
	Sum uint32 `json:"dep_sum"`
}
