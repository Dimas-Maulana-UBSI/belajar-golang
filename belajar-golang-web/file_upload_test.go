package belajar_golang_web

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//go:embed templates/*.gohtml
var Templ embed.FS

var Mytemplates = template.Must(template.ParseFS(Templ,"templates/*.gohtml"))
func UploadForm(w http.ResponseWriter, r *http.Request){
	Mytemplates.ExecuteTemplate(w,"form.gohtml",nil)
}

func Upload(w http.ResponseWriter,r *http.Request){
	file,fileHeader,err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination,err := os.Create("./resources/" + fileHeader.Filename )
	if err != nil {
		panic(err)
	}
	_,err = io.Copy(fileDestination,file)
	if err != nil {
		panic(err)
	}
	Mytemplates.ExecuteTemplate(w,"succes.gohtml",map[string]interface{}{
		"File" : "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/form",UploadForm)
	mux.HandleFunc("/upload",Upload)
	mux.Handle("/static/",http.StripPrefix("/static",http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	
}
//go:embed resources/smote.png
var UploadFileTest []byte
func TestUploadForm2(t *testing.T){
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file,_ := writer.CreateFormFile("file","contoh.png")
	file.Write(UploadFileTest)
	writer.Close()


	request := httptest.NewRequest(http.MethodPost,"localhost:8080/upload",body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder,request)

	respon,_ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(respon))
}