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
		"ToDoUpdate",
		"PUT",
		"/v1/todo",
		ToDoUpdate,
	},
	Route{
		"ToDoById",
		"GET",
		"/v1/todo/{id}",
		ToDoByID,
	},
	Route{
		"ToDoRemoveByID",
		"DELETE",
		"/v1/todo/{id}",
		ToDoRemoveByID,
	},
}
