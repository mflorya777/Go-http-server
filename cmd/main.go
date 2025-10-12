package main

import (
	"fmt"
	"http-server/configs"
	"http-server/internal/hello"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
