package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JeromeDesseaux/goInit/src/controllers"
)

func main() {
	http.HandleFunc("/hello", controllers.Hello)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
