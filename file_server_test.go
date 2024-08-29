package goweb

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {

	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8088",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

//go:embed resources
var resources embed.FS

func TestEmbedFileServer(t *testing.T) {

	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8088",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("terjadi error : ")
		panic(err)
	}

}
