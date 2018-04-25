package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

var tdll = ToDoList{nil, nil, 0}

func main() {

	tdll.AppendToDo(MakeToDo("First", "Testing"))

	router := mux.NewRouter()
	router.HandleFunc("/v1/todo", ToDoIndex).Methods("GET")
	router.HandleFunc("/v1/todo", ToDoCreate).Methods("POST")
	router.HandleFunc("/v1/todo/{id}", ToDoByID).Methods("GET")

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
	todos := tdll.GetArray(tdll.Size)
	json.NewEncoder(w).Encode(todos)
}

func ToDoCreate(w http.ResponseWriter, r *http.Request) {
	var todo ToDo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	newToDo := tdll.CreateToDo(todo)
	tdll.AppendToDo(newToDo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func ToDoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.FromString(vars["id"])
	todo, err := tdll.GetToDoByID(id)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func ToDoRemoveById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.FromString(vars["id"])
	tdll.RemoveToDoByID(id)
}
