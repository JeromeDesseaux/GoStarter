package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JeromeDesseaux/goInit/src/config"
	"github.com/JeromeDesseaux/goInit/src/controllers"
)

func main() {
	http.HandleFunc("/hello", controllers.Hello)

	config := config.GetConfig()

	fmt.Printf("Starting server at port %d\n", config.ApiPort)
	port := fmt.Sprintf(":%d", config.ApiPort)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
