package main

import (
	"net/http"

	"github.com/MachadoMichael/morpheus-proxy/config"
	"github.com/MachadoMichael/morpheus-proxy/handler"
)

func main() {

	err := config.Init()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler.HandleRequestAndRedirect)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
