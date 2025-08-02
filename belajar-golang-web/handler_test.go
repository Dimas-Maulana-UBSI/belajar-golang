package belajar_golang_web
import (
	"net/http"
	// "net/http/test"
	"testing"
	"fmt"
)
func TestHandler(t *testing.T){
	var handler http.HandlerFunc =  func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w,"halo dunia")
	}

	server := http.Server{
		Addr : "localhost:8080",
		Handler : handler,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}

}

func TestServeMux(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprint(w,"halo dunia")
		fmt.Fprintln(w,r.Method)
		fmt.Fprintln(w,r.Header)
		fmt.Fprintln(w,r.Body)
	})
	mux.HandleFunc("/hi",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprint(w,"hi")
		fmt.Fprintln(w,r.Method)
		fmt.Fprintln(w,r.Header)
		fmt.Fprintln(w,r.Body)
	})
	mux.HandleFunc("/images/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprint(w,"hi")
		fmt.Fprintln(w,r.Method)
		fmt.Fprintln(w,r.Header)
		fmt.Fprintln(w,r.Body)
	})
	mux.HandleFunc("/images/thumbnails",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprint(w,"hi")
		fmt.Fprintln(w,r.Method)
		fmt.Fprintln(w,r.Header)
		fmt.Fprintln(w,r.Body)
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
