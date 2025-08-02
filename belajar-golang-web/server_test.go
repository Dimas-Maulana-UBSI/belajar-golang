package belajar_golang_web

import (
	"testing"
	"net/http"
)

func TestServer(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
	}

	err:= server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}