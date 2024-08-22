package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	first_name := request.URL.Query().Get("first_name")
	last_name := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", first_name, last_name)
}

func TestQueryParam(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8880?first_name=moh&last_name=syaugi", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	read, _ := io.ReadAll(response.Body)
	bodyString := string(read)

	fmt.Println(bodyString)

}

func MultipleParamValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprintf(writer, strings.Join(names, " "))
}

func TestMultipleParamValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8880?name=moh&name=syaugi", nil)
	recorder := httptest.NewRecorder()

	MultipleParamValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
