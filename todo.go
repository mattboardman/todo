package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

// ToDo is a node in a linked list
// It contains self-descriptive properties
type ToDo struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Order       int       `json:"order,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	StartedOn   time.Time `json:"started,omitempty"`
	CompletedOn time.Time `json:"completed,omitempty"`
	Completed   bool
	Next        *ToDo
	Prev        *ToDo
}

// ToDoList is a Linked List data structure
// It contains the Head and Tail nodes as well as size
type ToDoList struct {
	Head *ToDo
	Tail *ToDo
	Size int
}

// The layout being used for displaying time
const layout = "Jan 2, 2006 at 3:04pm (CST)"

// ToString prints the editable properties of the ToDo struct
func (t *ToDo) ToString() string {
	properties := fmt.Sprintf(
		"Order:%d\t"+
			"Title:%s\t"+
			"Description:%s\t"+
			"StartedOn:%s\t"+
			"CompletedOn:%s\t"+
			"Completed: %t\t",
		t.Order,
		t.Title,
		t.Description,
		t.StartedOn.UTC().Format(layout),
		t.CompletedOn.UTC().Format(layout))
	return properties
}

// MakeToDo creates a new ToDo item
// Default properties are:
// ID: UUID, StartedOn: Current Time, CompletedOn: Beginning of Time
// Completed: False
// It returns a new ToDo struct
func MakeToDo(title, description string) *ToDo {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal("Fatal error creating UUID")
	}

	todo := ToDo{*id, 0, title, description, time.Now(), time.Time{}, false, nil, nil}
	return &todo
}

// UpdateToDoEntity is a ToDoList method that updates an existing
// ToDo item with a new Title, Description, Order, Completed
// returns an error from calling GetToDoByID() method
func (tdl *ToDoList) UpdateToDoEntity(newToDo ToDo) error {
	oldToDo, err := tdl.GetToDoByID(newToDo.ID)
	if err != nil {
		return err
	}

	oldToDo.Title = newToDo.Title
	oldToDo.Description = newToDo.Description
	if newToDo.Completed {
		oldToDo.CompletedOn = time.Now()
		oldToDo.Completed = true
	}

	if oldToDo.Order != newToDo.Order {
		oldToDo.Order = newToDo.Order
		tdl.insertToDo(oldToDo)
	}

	return nil
}

// AppendToDo is a ToDoList method that adds a new ToDo
// item to the end of the list. If the list is empty
// it points the List head and tail at the new ToDo item
func (tdl *ToDoList) AppendToDo(newToDo *ToDo) {
	if tdl.Head == nil {
		tdl.Head = newToDo
		tdl.Tail = newToDo
		newToDo.Order = 1
	} else {
		tdl.Tail.Next = newToDo
		newToDo.Prev = tdl.Tail
		tdl.Tail = newToDo
		newToDo.Order = newToDo.Prev.Order + 1
	}

	tdl.Size++
}

// insertToDo is a ToDoList method that inserts a ToDo item
// at it's current Order value in the list. It shifts values
// >= to the right and increments their order number
func (tdl *ToDoList) insertToDo(todo *ToDo) {
	var currentNode *ToDo

	if todo.Order == 1 {
		tdl.Head.Prev = todo
		todo.Next = tdl.Head
		tdl.Head = todo
		currentNode = todo.Next
	} else {
		currentNode = tdl.Head
	}

	for currentNode.Next != nil {
		if currentNode.Order == todo.Order {
			todo.Next = currentNode
			todo.Prev = currentNode.Prev
			todo.Prev.Next = todo
			currentNode.Order++
		} else if currentNode.Order > todo.Order {
			currentNode.Order++
		}
		currentNode = currentNode.Next
	}
	currentNode.Order++
	tdl.Size++
}

// RemoveToDoByID is a ToDoList method that removes
// a ToDo item associated with the input parameter id
// it returns an error from the called GetToDoByID() method
func (tdl *ToDoList) RemoveToDoByID(id uuid.UUID) error {
	todo, err := tdl.GetToDoByID(id)
	if err != nil {
		return err
	}

	if tdl.Head == todo {
		tdl.Head = todo.Next
		todo.Next.Prev = nil
	} else if tdl.Tail == todo {
		tdl.Tail = todo.Prev
		tdl.Tail.Next = nil
		tdl.Size--
		return nil
	} else {
		todo.Prev.Next = todo.Next
		todo.Next.Prev = todo.Prev
	}

	currentNode := todo
	for currentNode.Next != nil {
		currentNode.Order--
		currentNode = currentNode.Next
	}
	currentNode.Order--
	tdl.Size--
	return nil
}

// GetToDoByID is a ToDoList method that traverses the list
// until it finds a ToDo item that matches
// returns a ToDo item and an error if it fails to find a match
func (tdl *ToDoList) GetToDoByID(id uuid.UUID) (*ToDo, error) {
	if tdl.Head == nil {
		return nil, errors.New("No To-Do items added yet")
	}

	currentNode := tdl.Head
	for currentNode.Next != nil {
		if currentNode.ID == id {
			return currentNode, nil
		}
		currentNode = currentNode.Next
	}

	if currentNode.ID == id {
		return currentNode, nil
	}

	return nil, errors.New("Could not find To-Do item")
}

// PrintToDoList is a ToDoList method that prints out
// all of the nodes in the list
func (tdl *ToDoList) PrintToDoList() {
	currentNode := tdl.Head
	for currentNode != nil {
		fmt.Println(currentNode.ToString())
		currentNode = currentNode.Next
	}
}

// IsEmpty is a ToDoList method that determines if
// the list has been appended to or not
func (tdl *ToDoList) IsEmpty() bool {
	if tdl.Head == nil {
		return true
	}
	return false
}
