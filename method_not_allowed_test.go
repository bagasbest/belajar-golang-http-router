package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func NotAllowed(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Method Request Lu Salah Cuk")
}

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(NotAllowed)
	router.POST("/", SayHello)

	/// harus pake slash / di ujung
	request := httptest.NewRequest("POST", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Method Request Lu Salah Cuk", string(body))
	//fmt.Println(string(body))
}
