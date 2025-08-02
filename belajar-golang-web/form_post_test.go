package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	nama_depan := r.PostForm.Get("nama_depan")
	nama_belakang := r.PostForm.Get("nama_belakang")

	fmt.Fprintf(w,"halo %s %s",nama_depan,nama_belakang)
}

func TestFormPost(t *testing.T){
	requestBody := strings.NewReader("nama_depan=Dimas&nama_belakang=Maulana")
	request := httptest.NewRequest(http.MethodPost,"localhost:8080",requestBody)
	request.Header.Add("Content-Type","application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	FormPost(recorder,request)
	body,_ := io.ReadAll(recorder.Body)

	fmt.Println(string(body))
}