package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!\n"))
	})
	//FIXME: from configuration
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
