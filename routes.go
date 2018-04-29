package main

import (
	"net/http"
)

// Route struct contains
// Route Name
// REST request
// Path
// Handler Function
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is an array of Routes
type Routes []Route

var routes = Routes{
	Route{
		"ToDoIndex",
		"GET",
		"/v1/todo",
		ToDoIndex,
	},
	Route{
		"ToDoReverseIndex",
		"GET",
		"/v1/todo/reverse",
		ToDoReverseIndex,
	},
	Route{
		"ToDoCompleted",
		"GET",
		"/v1/todo/completed",
		ToDoCompleted,
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
	Route{
		"ToDoSearchByTitle",
		"GET",
		"/v1/todo/search/{title}",
		ToDoSearchByTitle,
	},
	Route{
		"ToDoImprovedSearchByTitle",
		"GET",
		"/v2/todo/search/{title}",
		ToDoImprovedSearchByTitle,
	},
}
