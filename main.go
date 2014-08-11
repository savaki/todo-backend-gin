package main

import (
	"net/http"
	"os"
)

func main() {
	todo := Todo{}
	routes := makeRoutes(todo)
	http.ListenAndServe(":"+os.Getenv("PORT"), routes)
}
