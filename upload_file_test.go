package goweb

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	// request.ParseMultipartForm(100 << 20) // untuk merubah size file
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}

	fmt.Println(fileHeader.Filename)

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	// copy file ke file destinatio
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := request.PostFormValue("name")

	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})

}

func TestUploadForm(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    ":8088",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
