package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {

	cookie := new(http.Cookie) //membuat cookie
	cookie.Name = "X-PZN-name" //set nama(key) cookie
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprintf(writer, "success create cookie")

}

func GetCookie(writer http.ResponseWriter, request *http.Request) {

	cookie, err := request.Cookie("X-PZN-name")
	if err != nil {
		fmt.Fprintf(writer, "tidak ada cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "hello this cookie value is %s", name)
	}

}

func TestCookie(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8089",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestSetCookie(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=ogi", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("cookie %s:%s \n", cookie.Name, cookie.Value)
	}

}

func TestGetCookie(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=ogi", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-name"
	cookie.Value = "OGI"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}
