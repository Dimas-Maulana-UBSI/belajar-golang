package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HaloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"halo dunia")
}

func TestHttp(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080/halo",nil)
	recorder := httptest.NewRecorder()

	HaloHandler(recorder,request)

	respone := recorder.Result()

	body,_ := io.ReadAll(respone.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

