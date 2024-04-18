package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	/// BEDANYA http.ServeMux dengan httprouter adalah kalo http router bisa kita tentukan methodnya
	/// jadi, request api hanaya bisa di request dengan method tertentu saja contohnya GET
	/// karena kalo serve mux, mau get, post, dll, itu masih bisa di request
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello HttpRouter")
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
