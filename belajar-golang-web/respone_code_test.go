package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponeCode(w http.ResponseWriter,r *http.Request){
	nama := r.URL.Query().Get("nama")
	if nama == ""{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"salah kirim")
	}else{
		fmt.Fprintf(w,"halo %s",nama)
	}
}

func TestResponeCode(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"localhost:8080?nama=Dimas",nil)
	recorder := httptest.NewRecorder()

	ResponeCode(recorder,request)
	respone := recorder.Result()

	body,_:= io.ReadAll(respone.Body)

	fmt.Println(respone.Status)
	fmt.Println(respone.StatusCode)
	fmt.Println(string(body))
}