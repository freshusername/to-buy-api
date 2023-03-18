package api

import (
	"net/http"

	"github.com/freshusername/to-buy-api/types"
)

func HandleGetAccount(w http.ResponseWriter, r *http.Request) {
	var account types.Account
	_ = account
	w.Write([]byte("Ok from account :)"))
}
