package main

import (
	"github.com/annawinkler/BuildingRestfulAPIswithGo/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Starting server on localhost:11111")
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
