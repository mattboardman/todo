package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	uuid "github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var tdll = ToDoList{nil, nil, 0}
var bstID = Tree{nil}
var bstTitle = Tree{nil}

func main() {
	first := MakeToDo("First", "Testing")
	id, _ := uuid.Parse("44f52c01-ddf2-459d-be19-44c057719f74")
	second := MakeToDo("Second", "Testing")

	second.ID = id
	tdll.AppendToDo(first)
	tdll.AppendToDo(second)

	bstID.InsertByID(first)
	bstID.InsertByID(second)

	bstTitle.InsertByString(first)
	bstTitle.InsertByString(second)

	addToDos(30000000) // 30,000,000
	log.Println("Done Loading Test Data")

	router := NewRouter()

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
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
	if todos := tdll.GetArray(100); todos != nil {
		json.NewEncoder(w).Encode(todos)
	}
}

func ToDoCreate(w http.ResponseWriter, r *http.Request) {
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
		bstID.InsertByID(newToDo)
		bstTitle.InsertByString(newToDo)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newToDo); err != nil {
			panic(err)
		}
	}
}

func ToDoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
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
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	tdll.RemoveToDoByID(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ToDoUpdate(w http.ResponseWriter, r *http.Request) {
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

	old, _ := tdll.GetToDoByID(todo.ID)
	bstTitle.DeleteByString(old.Title)
	tdll.UpdateToDoEntity(todo)
	bstTitle.InsertByString(&todo)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func ToDoSearchByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title, _ := vars["title"]
	//todo := bstTitle.FindByString(title)
	todo, _ := tdll.GetToDoByTitle(title)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func ToDoImprovedSearchByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title, _ := vars["title"]
	todo := bstTitle.FindByString(title)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func addToDos(num int) {
	bottom := "zzzzzz"
	for i := 0; i < num; i++ {
		todo := MakeToDo(RandStringRunes(10), bottom)
		tdll.AppendToDo(todo)
	}
	todo := MakeToDo(bottom, bottom)
	tdll.AppendToDo(todo)
	bstID.InsertByID(todo)
	bstTitle.InsertByString(todo)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
