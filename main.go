package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

var tdll = ToDoList{nil, nil, 0}

func main() {

	tdll.AppendToDo(MakeToDo("First", "Testing"))
	id, _ := uuid.FromString("44f52c01-ddf2-459d-be19-44c057719f74")

	test := MakeToDo("Second", "Testing")
	test.ID = id
	tdll.AppendToDo(test)

	router := NewRouter()

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{
		"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
		wg.Done()
	}()

	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/", fs)

	log.Println("Listening...")

	wg.Add(1)
	go func() {
		http.ListenAndServe(":1080", nil)
		wg.Done()
	}()
	wg.Wait()
}

func ToDoIndex(w http.ResponseWriter, r *http.Request) {
	if todos := tdll.GetArray(tdll.Size); todos != nil {
		json.NewEncoder(w).Encode(todos)
	}
}

func ToDoCreate(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var todo ToDo

	params := strings.Split(r.URL.RawQuery, "?")
	title := strings.Split(params[0], "=")
	description := strings.Split(params[1], "=")
	if todo.Title = title[1]; todo.Title == "" {
		w.WriteHeader(422)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		todo.Description = description[1]
		newToDo := tdll.CreateToDo(todo)
		tdll.AppendToDo(newToDo)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newToDo); err != nil {
			panic(err)
		}
	}
}

func ToDoByID(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	id, _ := uuid.FromString(vars["id"])
	todo, err := tdll.GetToDoByID(id)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func ToDoRemoveByID(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	id, _ := uuid.FromString(vars["id"])
	tdll.RemoveToDoByID(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ToDoUpdate(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

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

	tdll.UpdateToDoEntity(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, HEAD")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
