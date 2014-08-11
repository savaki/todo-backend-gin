package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func makeRoutes(todo Todo) http.Handler {
	ok := func(c *gin.Context) {
		c.String(200, "")
	}

	cors := func(c *gin.Context) {
		c.Writer.Header().Add("access-control-allow-origin", "*")
		c.Writer.Header().Add("access-control-allow-headers", "accept, content-type")
		c.Writer.Header().Add("access-control-allow-methods", "GET,HEAD,POST,DELETE,OPTIONS,PUT,PATCH")
	}

	routes := gin.Default()
	routes.Use(cors)

	routes.OPTIONS("/todos", ok)

	routes.OPTIONS("/todos/:id", ok)

	routes.GET("/todos", func(c *gin.Context) {
		c.JSON(200, todo.All())
	})

	routes.GET("/todos/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.JSON(200, todo.Find(id))
	})

	routes.POST("/todos", func(c *gin.Context) {
		template := TodoItem{}
		if c.EnsureBody(&template) {
			fqdn := func(path string) string {
				return "http://todo-backend-golang.herokuapp.com" + path
			}
			item := todo.Create(template, fqdn)
			c.Writer.Header().Add("Location", "/todos/"+item.Id)
			c.JSON(201, item)
		}
	})

	routes.PATCH("/todos/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		if item := todo.Find(id); item != nil {
			c.EnsureBody(&item)
			c.JSON(200, item)
		}
	})

	routes.DELETE("/todos", func(c *gin.Context) {
		c.JSON(200, todo.DeleteAll())
	})

	routes.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.JSON(200, todo.Delete(id))
	})

	return routes
}
