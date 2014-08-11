package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	todo := Todo{}
	routes := makeRoutes(todo)
	http.ListenAndServe(":"+os.Getenv("PORT"), routes)
}
