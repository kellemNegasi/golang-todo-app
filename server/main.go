package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kellemnegasi/todo-app/server/router"
)

func main() {
	fmt.Println("trying")
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
