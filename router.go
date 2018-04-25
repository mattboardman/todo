package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"Get",
		"/",
		Index,
	},
	Route{
		"ToDoIndex",
		"GET",
		"/v1/todo",
		ToDoIndex,
	},
	Route{
		"ToDoCreate",
		"POST",
		"/v1/todo",
		ToDoCreate,
	},
	Route{
		"ToDoById",
		"GET",
		"/v1/todo/{id}",
		ToDoByID,
	},
}
