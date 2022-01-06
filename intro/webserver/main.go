package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello",func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello world!"))
	} )

	if err := http.ListenAndServe(":8080",nil); err != nil {
		panic(err)
	}

}