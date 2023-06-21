package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World!!!"
	fmt.Fprint(w, msg)
}

func main() {
	http.HandleFunc("/", helloHandle)
	fmt.Println("http://localhost:8080 で起動中...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
