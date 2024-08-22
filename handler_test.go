package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "hello world")
		if err != nil {
			fmt.Println("error write in handler")
			panic(err)
		}
	}

	server := http.Server{
		Addr:    "localhost:8880",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "hello world")
	})

	mux.HandleFunc("/hi", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "hi world")
	})

	server := http.Server{
		Addr:    "localhost:8880",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("server ready on port 8880....")
	}

}
