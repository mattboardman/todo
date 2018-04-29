package main

import (
	"encoding/json"
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

// ToDoList
var tdll = ToDoList{nil, nil, 0}

// ID based binary search tree
var bstID = Tree{nil}

// Title based binary search tree
var bstTitle = Tree{nil}

func main() {
	id, _ := uuid.Parse("44f52c01-ddf2-459d-be19-44c057719f74")
	first := MakeToDo("First", "Testing")

	first.ID = id
	tdll.AppendToDo(&first)

	bstID.InsertByID(&first)

	bstTitle.InsertByString(&first)

	addToDos(1000) // 1,000
	log.Printf("Done Loading Test Data: %d records", tdll.Size)

	router := NewRouter()

	// Allow all requests for all routes - Channge if not in DEV environment
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{
		"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	wg := &sync.WaitGroup{}

	// API is setup
	wg.Add(1)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
		wg.Done()
	}()

	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/", fs)

	log.Println("Listening...")

	// Frontend is setup
	wg.Add(1)
	go func() {
		http.ListenAndServe(":1080", nil)
		wg.Done()
	}()
	wg.Wait()
}

// ToDoIndex gets the first 100 items in the repository
func ToDoIndex(w http.ResponseWriter, r *http.Request) {
	if todos := tdll.GetArray(0, 100); todos != nil {
		json.NewEncoder(w).Encode(todos)
	}
}

// ToDoReverseIndex gets the last 100 records in the repository
func ToDoReverseIndex(w http.ResponseWriter, r *http.Request) {
	if todos := tdll.GetReverseArray(tdll.Size - 100); todos != nil {
		json.NewEncoder(w).Encode(todos)
	}
}

// ToDoCompleted gets the first 10 completed ToDos
func ToDoCompleted(w http.ResponseWriter, r *http.Request) {
	if todos := tdll.GetAllCompleted(10); todos != nil {
		json.NewEncoder(w).Encode(todos)
	}
}

// ToDoCreate adds a new ToDo item to:
// 1. ToDoList
// 2. id bst
// 3. string bst
// Create Time: O(1) + O(2log(n))
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

// ToDoByID returns a ToDo item by its uuid
// Retrieval Time: O(log(n))
func ToDoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	todo := bstID.FindByID(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

// ToDoRemoveByID removes a ToDo from the:
// 1. ToDoList
// 2. id bst
// 3. string bst
// Remove Time: O(1) + O(2log(n))
func ToDoRemoveByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	todo, _ := tdll.GetToDoByID(id)
	tdll.RemoveToDoByID(todo.ID)
	bstID.DeleteByID(todo.ID)
	bstTitle.DeleteByString(todo.Title)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// ToDoUpdate updates an existing ToDo item
// 1. The ToDo is updated directly
// 2. The ToDo item is removed from the string bst
// 3. The ToDo is re-added to the string bst
// Update Time: O(1) + O(2log(n))
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

// ToDoSearchByTitle searches the repository using a linked
// list traverse in single order
// Search Time: O(n)
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

// ToDoImprovedSearchByTitle searches the repository using the
// string binary search tree created at run time
// Search Time: O(log(n))
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
		item := MakeToDo(RandStringRunes(10), bottom)
		tdll.AppendToDo(&item)
		bstID.InsertByID(&item)
		bstTitle.InsertByString(&item)
	}
	todo := MakeToDo(bottom, bottom)
	tdll.AppendToDo(&todo)
	bstID.InsertByID(&todo)
	bstTitle.InsertByString(&todo)
}

// Seed the random number generator on startup
func init() {
	rand.Seed(time.Now().UnixNano())
}

// letters to seed the string generator
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes generates strings of 'n' length
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
