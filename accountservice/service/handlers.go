package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/caominhtri91/micro/accountservice/dbclient"
	"github.com/gorilla/mux"
)

// DBClient variable
var DBClient dbclient.IBoltClient

// GetAccount handler read the id and give the account
func GetAccount(w http.ResponseWriter, r *http.Request) {
	// Read the 'accountID' path parameter from the mux map
	var accountID = mux.Vars(r)["accountID"]

	// Read the account struct BoltDB
	account, err := DBClient.QueryAccount(accountID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
