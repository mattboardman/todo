package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

type ToDo struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Order       int       `json: "order,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Started     time.Time `json:"started,omitempty"`
	Completed   time.Time `json:"completed,omitempty"`
	Next        *ToDo
	Prev        *ToDo
}

type ToDoList struct {
	Head *ToDo
	Tail *ToDo
	Size int
}

const layout = "Jan 2, 2006 at 3:04pm (CST)"

func (t *ToDo) ToString() string {
	properties := fmt.Sprintf(
		"Order:%d\t"+
			"Title:%s\t"+
			"Description:%s\t"+
			"Started:%s\t"+
			"Completed:%s\t",
		t.Order,
		t.Title,
		t.Description,
		t.Started.UTC().Format(layout),
		t.Completed.UTC().Format(layout))
	return properties
}

func MakeToDo(title, description string) *ToDo {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal("Fatal error creating UUID")
	}

	todo := ToDo{*id, 0, title, description, time.Now(), time.Time{}, nil, nil}
	return &todo
}

func (tdl *ToDoList) UpdateToDoEntity(newToDo ToDo) error {
	oldToDo, err := tdl.GetToDoByID(newToDo.ID)
	if err != nil {
		return err
	}

	oldToDo.Title = newToDo.Title
	oldToDo.Description = newToDo.Description
	if oldToDo.Order != newToDo.Order {
		oldToDo.Order = newToDo.Order
		tdl.InsertToDo(oldToDo)
	}
	return nil
}

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

func (tdl *ToDoList) InsertToDo(todo *ToDo) {
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

func (tdl *ToDoList) PrintToDoList() {
	currentNode := tdl.Head
	for currentNode != nil {
		fmt.Println(currentNode.ToString())
		currentNode = currentNode.Next
	}
}

func (tdl *ToDoList) IsEmpty() bool {
	if tdl.Head == nil {
		return true
	}
	return false
}
