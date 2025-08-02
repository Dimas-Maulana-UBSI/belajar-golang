package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequesHeader(w http.ResponseWriter,r *http.Request){
	header := r.Header.Get("content-type")
	fmt.Fprint(w,header)
}

func TestRequestHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodPost,"localhost:8080/post",nil)
	request.Header.Add("content-type","aplication/json")

	recorder := httptest.NewRecorder()

	RequesHeader(recorder,request)
	respone := recorder.Result()
	body,_:= io.ReadAll(respone.Body)
	fmt.Println(string(body))

}

func ResponeHeader(w http.ResponseWriter, r *http.Request){
	w.Header().Add("x-powered-by","dimas maulana")
	fmt.Fprint(w,"ok")
}

func TestResponeHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodPost,"localhost:8080/post",nil)
	request.Header.Add("content-type","aplication/json")

	recorder := httptest.NewRecorder()

	ResponeHeader(recorder,request)
	respone := recorder.Result()
	body,_:= io.ReadAll(respone.Body)
	fmt.Println(string(body))
	fmt.Println(respone.Header.Get("x-powered-by"))

}