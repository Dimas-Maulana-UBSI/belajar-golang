package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
	_"embed"
)

func ServeFile(w http.ResponseWriter,r *http.Request){
	if r.URL.Query().Get("nama") != ""{
		http.ServeFile(w,r,"./resources/dimas.html")//serve file tidak bisa menggunakan golang embed
	}else{
		http.ServeFile(w,r,"./resources/index.html")

	}

}

func TestServeFile(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


//go:embed resources/dimas.html
var dimas string

//go:embed resources/index.html
var index string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request){
	if r.URL.Query().Get("nama") != ""{
		fmt.Fprint(w,dimas)
	}else{
		fmt.Fprint(w,index)
	}
}

func TestServeFileEmbed(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}