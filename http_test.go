package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHttpTest(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8880", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	read, _ := io.ReadAll(response.Body)
	bodyString := string(read)

	fmt.Println(bodyString)
}
