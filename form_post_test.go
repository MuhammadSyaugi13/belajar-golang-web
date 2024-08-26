package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writter http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()

	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	fmt.Fprintf(writter, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {

	requestBody := strings.NewReader("first_name=moh&last_name=ogi")
	request := httptest.NewRequest(http.MethodPost, "http://localhost/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
