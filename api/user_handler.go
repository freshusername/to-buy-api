package api

import (
	"net/http"

	"github.com/freshusername/to-buy-api/types"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	var user types.User
	_ = user

	w.Write([]byte("Ok from user"))
}
