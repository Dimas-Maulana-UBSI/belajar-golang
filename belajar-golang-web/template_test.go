package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS

type data struct {
	Nama string
	Id int
	Alamat interface{}
}
func TemplateEmbed(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFS(templates,"templates/*.gohtml"))
	t.ExecuteTemplate(w,"index.gohtml",data{
		Nama: "dimas maulana",
		Id: 1,
		Alamat: "kebon kelapa",
	})
}

func TestTemplateEmbed(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder,request)

	result,_ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(result))


}

func TemplateAction(w http.ResponseWriter,r *http.Request){
	t := template.Must(template.ParseFS(templates,"templates/*.gohtml"))
	t.ExecuteTemplate(w,"action.gohtml",data{
		Nama: "kambing",
		Id: 5,
		Alamat: map[string]string{
			"Jalan": "jl kambing",
			"Kecamatan": "peternakan",
		},
	})
}

func TestTemplateAction(t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"localhost:8080",nil)
	recorder := httptest.NewRecorder()
	TemplateAction(recorder,request)
	body,err:= io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}