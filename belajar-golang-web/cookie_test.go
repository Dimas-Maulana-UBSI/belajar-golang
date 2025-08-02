package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request){
	cookie := new(http.Cookie)
	cookie.Name = "dimas-maulana"
	cookie.Value = r.URL.Query().Get("nama")
	cookie.Path = "/"

	http.SetCookie(w,cookie)
	fmt.Fprintf(w,"berhasil menambahkan cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request){
	cookie,err := r.Cookie("dimas-maulana")
	if err != nil {
		fmt.Fprint(w,"tidak ada cookie")
	}else{
		nama := cookie.Value
		fmt.Fprintf(w,"halo %s",nama)
	}
}

func TestSetCookie(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"localhost:8080?nama=Dimas",nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder,request)

	cookies := recorder.Result().Cookies()

	for _,cookie := range cookies{
		fmt.Printf("cookie %s %s \n",cookie.Name,cookie.Value)
	}
}

func TestGetCookie(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"localhost:8080",nil)
	cookie := new(http.Cookie)

	cookie.Name = "dimas-maulana"
	cookie.Value = "dimas"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder,request)
	body,_ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}