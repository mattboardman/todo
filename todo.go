package main

import (
	"errors"
	"fmt"
	"time"

	uuid "github.com/google/uuid"
)

// ToDo is a node in a linked list
// It contains self-descriptive properties
type ToDo struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	StartedOn   time.Time `json:"started,omitempty"`
	CompletedOn time.Time `json:"completed,omitempty"`
	IsCompleted bool      `json:"iscompleted"`
	Next        *ToDo     `json:"-"`
	Prev        *ToDo     `json:"-"`
}

// ToDoList is a Linked List data structure
// It contains the Head and Tail nodes as well as size
type ToDoList struct {
	Head *ToDo
	Tail *ToDo
	Size int
}

// ToDos is an array of ToDo items
type ToDos []ToDo

// The layout being used for displaying time
const layout = "Jan 2, 2006 at 3:04pm (CST)"

// ToString prints the editable properties of the ToDo struct
func (t *ToDo) ToString() string {
	properties := fmt.Sprintf(
		"Title:%s\t"+
			"Description:%s\t"+
			"StartedOn:%s\t"+
			"CompletedOn:%s\t"+
			"IsCompleted: %t\t",
		t.Title,
		t.Description,
		t.StartedOn.UTC().Format(layout),
		t.CompletedOn.UTC().Format(layout),
		t.IsCompleted)
	return properties
}

// MakeToDo creates a new ToDo item
// Default properties are:
// ID: UUID, StartedOn: Current Time, CompletedOn: Beginning of Time
// Completed: False
// It returns a new ToDo struct
func MakeToDo(title, description string) ToDo {
	id := uuid.New()

	todo := ToDo{id, title, description, time.Now(), time.Time{}, false, nil, nil}
	return todo
}

// GetArray gets an array of items from
// startIndex -> endIndex (inclusive, exclusive)
// If the index is invalid it will get the first item
func (tdl *ToDoList) GetArray(startIndex, endIndex int) ToDos {
	var todos ToDos
	if startIndex > tdl.Size {
		startIndex = 0
		endIndex = 1
		todos = make(ToDos, endIndex)

	} else {
		todos = make(ToDos, (endIndex - startIndex))
	}

	currentNode := tdl.Head
	if currentNode == nil {
		panic(errors.New("List has no items"))
	}

	count := 0
	for count < (endIndex-startIndex) && currentNode != nil {
		todos[count] = *currentNode
		count++
		currentNode = currentNode.Next
	}

	return todos
}

// GetReverseArray gets an array of items from
// End of ToDoList -> startIndex (inclusive, exclusive)
// If the index is invalid it will get the last item
func (tdl *ToDoList) GetReverseArray(startIndex int) ToDos {
	var todos ToDos
	if startIndex > tdl.Size {
		startIndex = tdl.Size - 1
		todos = make(ToDos, 1)
	} else {
		todos = make(ToDos, (tdl.Size - startIndex))
	}

	currentNode := tdl.Tail
	if currentNode == nil {
		panic(errors.New("List has no items"))
	}

	count := 0
	for count < (tdl.Size-startIndex) && currentNode != nil {
		todos[count] = *currentNode
		count++
		currentNode = currentNode.Prev
	}

	return todos
}

// GetAllCompleted grabs a ToDo[max] array
// of ToDo items completed in order of creation
func (tdl *ToDoList) GetAllCompleted(max int) ToDos {
	var todos ToDos
	if max > tdl.Size {
		max = tdl.Size
	}

	todos = make(ToDos, max)
	currentNode := tdl.Head
	if currentNode == nil {
		panic(errors.New("List has no items"))
	}

	count := 0
	numCompleted := 0
	for count < max && currentNode != nil {
		if currentNode.IsCompleted {
			todos[count] = *currentNode
			count++
			numCompleted++
		}
		currentNode = currentNode.Next
	}

	completedToDos := todos[:numCompleted]
	return completedToDos
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
	if newToDo.IsCompleted {
		oldToDo.CompletedOn = time.Now()
		oldToDo.IsCompleted = true
	} else if !newToDo.IsCompleted && oldToDo.IsCompleted {
		oldToDo.CompletedOn = time.Time{}
		oldToDo.IsCompleted = false
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
	} else {
		tdl.Tail.Next = newToDo
		newToDo.Prev = tdl.Tail
		tdl.Tail = newToDo
	}

	tdl.Size++
}

// CreateToDo returns a new ToDo item from an existing
// ToDo item
func (tdl *ToDoList) CreateToDo(newToDo ToDo) *ToDo {
	id := uuid.New()

	newToDo.ID = id
	newToDo.StartedOn = time.Now()
	newToDo.CompletedOn = time.Time{}
	newToDo.IsCompleted = false

	return &newToDo
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
		if tdl.Head.Next != nil {
			tdl.Head = todo.Next
			todo.Next.Prev = nil
		} else {
			tdl.Head = nil
			tdl.Tail = nil
		}
	} else if tdl.Tail == todo {
		tdl.Tail = todo.Prev
		tdl.Tail.Next = nil
	} else {
		todo.Prev.Next = todo.Next
		todo.Next.Prev = todo.Prev
	}

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
	for currentNode != nil {
		if currentNode.ID == id {
			return currentNode, nil
		}
		currentNode = currentNode.Next
	}

	return nil, errors.New("Could not find To-Do item")
}

// GetToDoByTitle is a ToDoList method that traverses the list
// until it finds a ToDo item that matches
// returns a ToDo item and an error if it fails to find a match
func (tdl *ToDoList) GetToDoByTitle(value string) (*ToDo, error) {
	if tdl.Head == nil {
		return nil, errors.New("No To-Do items added yet")
	}

	currentNode := tdl.Head
	for currentNode != nil {
		time.Sleep(10 * time.Millisecond)
		if currentNode.Title == value {
			return currentNode, nil
		}
		currentNode = currentNode.Next
	}

	return nil, errors.New("Could not find To-Do item")
}

// resets a list to zero values
func (tdl *ToDoList) clearList() {
	tdl.Head = nil
	tdl.Tail = nil
	tdl.Size = 0
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
