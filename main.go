package main

import (
	"net/http"
	"time"

	"github.com/nu7hatch/gouuid"
)

func main() {

	/* 	router := mux.NewRouter()
	   	router.HandleFunc("/todo", GetToDoList).Methods("GET")
	   	router.HandleFunc("/todo/{id}", GetToDo).Methods("GET")
	   	router.HandleFunc("/todo/{id}", CreateToDo).Methods("POST")
	   	router.HandleFunc("/todo/{id}", UpdateToDo).Methods("PUT")
	   	router.HandleFunc("/todo/{id}", DeleteToDo).Methods("DELETE")
	   	log.Fatal(http.ListenAndServe(":8000", router)) */

	tdl := ToDoList{nil, nil, 0}

	tdl.AppendToDo(MakeToDo("First", "Testing"))
	tdl.AppendToDo(MakeToDo("Second", "Out"))
	tdl.AppendToDo(MakeToDo("Third", "Order"))
	var id4, _ = uuid.NewV4()
	var t4 = ToDo{*id4, 2, "Fourth", "Testing", time.Now(), time.Time{}, nil, nil}
	tdl.InsertToDo(&t4)
	tdl.PrintToDoList()

}

func GetToDoList(w http.ResponseWriter, r *http.Request) {}
func GetToDo(w http.ResponseWriter, r *http.Request)     {}
func CreateToDo(w http.ResponseWriter, r *http.Request)  {}
func UpdateToDo(w http.ResponseWriter, r *http.Request)  {}
func DeleteToDo(w http.ResponseWriter, r *http.Request)  {}
