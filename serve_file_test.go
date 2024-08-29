package goweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") == "" {
		http.ServeFile(writer, request, "./resources/notFound.html")
	} else {
		http.ServeFile(writer, request, "./resources/ok.html")
	}
}

func TestServeFile(t *testing.T) {

	server := http.Server{
		Addr:    ":8880",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

//go:embed resources/ok.html
var resourceOK string

//go:embed resources/notFound.html
var resourceNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") == "" {
		fmt.Fprint(writer, resourceNotFound)
	} else {
		fmt.Fprint(writer, resourceOK)
	}
}

func TestEmbedServeFile(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:8880",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
