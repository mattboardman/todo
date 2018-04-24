package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tdll = ToDoList{nil, nil, 0}

func main() {

	tdll.AppendToDo(MakeToDo("First", "Testing"))

	router := mux.NewRouter()
	router.HandleFunc("/v1/todo", ToDoIndex).Methods("GET")

	/* 	router.HandleFunc("/v1/todo/{id}", GetToDo).Methods("GET")
	   	router.HandleFunc("/v1/todo/{id}", CreateToDo).Methods("POST")
	   	router.HandleFunc("/v1/todo/{id}", UpdateToDo).Methods("PUT")
	   	router.HandleFunc("/v1/todo/{id}", DeleteToDo).Methods("DELETE") */
	log.Fatal(http.ListenAndServe(":8080", router))

	/*   	tdl.AppendToDo(MakeToDo("First", "Testing"))
	tdl.AppendToDo(MakeToDo("Second", "Out"))
	tdl.AppendToDo(MakeToDo("Third", "Order"))
	var id4, _ = uuid.NewV4()
	var t4 = ToDo{*id4, 2, "Fourth", "Testing", time.Now(), time.Time{}, false, nil, nil}
	tdl.insertToDo(&t4)
	tdl.PrintToDoList() */

}

func Index(w http.ResponseWriter, r *http.Request) {

}
func ToDoIndex(w http.ResponseWriter, r *http.Request) {
	todos := tdll.GetArray(2)
	json.NewEncoder(w).Encode(todos)
}
