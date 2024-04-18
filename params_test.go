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

func GetParams(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	text := "Product " + id
	fmt.Fprint(writer, text)
}

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", GetParams)

	/// harus pake slash / di ujung 
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/product/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))
	//fmt.Println(string(body))
}
