package main

import (
	"flag"
	"net/http"

	"github.com/freshusername/to-buy-api/api"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "todo")
	flag.Parse()

	http.HandleFunc("/user", api.HandleGetUser)
	http.HandleFunc("/account", api.HandleGetAccount)

	http.ListenAndServe(*listenAddr, nil)
}
