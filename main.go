package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			for _, s := range v {
				writer.Header().Add(k, s)
			}
		}

		writer.Header().Set("VERSION", os.Getenv("VERSION"))

		log.Printf("ip: %s, status: %d.", request.Host, 200)
	})

	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
