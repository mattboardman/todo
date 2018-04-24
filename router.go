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
}
