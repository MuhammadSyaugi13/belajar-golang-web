package goweb

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`

	// t, err := template.New("SIMPLE").Parse(templateText)

	// if err != nil {
	// 	panic(err)
	// }

	// t.ExecuteTemplate(writer, "SIMPLE", "Hello guys")

	t := template.Must(template.New("SIMPLE").Parse(templateText))
	t.ExecuteTemplate(writer, "SIMPLE", "Hello guys")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8088", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {

	// t := template.Must(template.ParseFiles("./templates/simple.gohtml")) //mengambil file simple.gohtml
	t := template.Must(template.ParseGlob("./templates/*.gohtml")) //mengambil file direktori templates yang memiliki extension .gohtml

	t.ExecuteTemplate(writer, "simple.gohtml", "hello html template")

}

func TestFileSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8088", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*dengan golang embed*/

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {

	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	t.ExecuteTemplate(writer, "simple.gohtml", "hello html template")

}

func TestEmbedTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8088", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/* ./ dengan golang embed*/

// dengan mux
func TestMUXFile(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8088", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		fmt.Println("error io.readAll")
		panic(err)
	}
	fmt.Println(string(body))

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, string(body))
		if err != nil {
			fmt.Println("error write in handler")
			panic(err)
		}
	}

	server := http.Server{
		Addr:    "localhost:8880",
		Handler: handler,
	}

	errs := server.ListenAndServe()
	if err != nil {
		panic(errs)
	}

}
