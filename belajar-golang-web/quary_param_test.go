package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHalo(w http.ResponseWriter, r *http.Request){
	nama := r.URL.Query().Get("nama")
	if nama == ""{
		fmt.Fprint(w,"halo")
	}else{
		fmt.Fprintf(w,"halo %s",nama)
	}
}

func TestQuartyParam(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"localhost:8080/halo?nama=Dimas",nil)
	recorder := httptest.NewRecorder()

	SayHalo(recorder,request)

	respone := recorder.Result()

	body,_ := io.ReadAll(respone.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func MultiQuary(w http.ResponseWriter, r *http.Request){
	namaDeapn := r.URL.Query().Get("nama_depan")
	namaBelakang := r.URL.Query().Get("nama_belakang")
	fmt.Fprintf(w,"halo nama saya %s %s",namaDeapn,namaBelakang)
}

func TestMultiQuaryParam(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"localhost:8080/halo?nama_depan=Dimas&nama_belakang=maulana",nil)
	recorder := httptest.NewRecorder()

	MultiQuary(recorder,request)

	respone := recorder.Result()

	body,_ := io.ReadAll(respone.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

func MultipleParamValues(w http.ResponseWriter, r *http.Request){
	quary := r.URL.Query()
	nama := quary["nama"]
	fmt.Fprint(w, strings.Join(nama, " "))
}

func TestMultiParamValues(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"localhost:8080/halo?nama=Dimas&nama=Maulana&nama=maul",nil)
	recorder := httptest.NewRecorder()

	MultipleParamValues(recorder,request)

	respone := recorder.Result()

	body,_ := io.ReadAll(respone.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}

