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

func NotFound(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Gak Ketemu")
}

func TestNotFound(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(NotFound)
	router.GET("/", SayHello)

	/// harus pake slash / di ujung
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/404", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak Ketemu", string(body))
	//fmt.Println(string(body))
}
