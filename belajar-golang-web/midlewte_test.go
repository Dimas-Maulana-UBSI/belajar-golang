package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware)ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Println("sebelum handler")
	middleware.Handler.ServeHTTP(w,r)
	fmt.Println("sesudah handler")
}
type HandleError struct {
	Handler http.Handler
}

func (middleware *HandleError)ServeHTTP(w http.ResponseWriter,r *http.Request){
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("terjadi error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w,"error : %s",err)
		}
	}()
	middleware.Handler.ServeHTTP(w,r)
}
func TestMiddleware(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handler executed")
		fmt.Fprint(w,"halo middleware")
	})
	mux.HandleFunc("/panic",func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("error executed")
		panic("yah panik")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandle := &HandleError{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: errorHandle,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}