package main

import (
	"testing"
	"time"

	uuid "github.com/google/uuid"
)

var tdl ToDoList

var id1 = uuid.New()
var id2 = uuid.New()
var id3 = uuid.New()
var id4 = uuid.New()
var t1 = ToDo{id1, "First", "Testing", time.Now(), time.Time{}, false, nil, nil}
var t2 = ToDo{id2, "Second", "Testing", time.Now(), time.Time{}, false, nil, nil}
var t3 = ToDo{id3, "Third", "Testing", time.Now(), time.Time{}, false, nil, nil}
var t4 = ToDo{id4, "Fourth", "Testing", time.Now(), time.Time{}, false, nil, nil}

func TestIsEmpty(t *testing.T) {
	tdl.clearList()

	if tdl.Head == nil && !tdl.IsEmpty() {
		t.Errorf("list should be empty")
	}

	tdl.AppendToDo(&t1)

	if tdl.Head != nil && tdl.IsEmpty() {
		t.Errorf("list should NOT be empty")
	}
}

func TestAppend(t *testing.T) {
	tdl.clearList()

	if !tdl.IsEmpty() {
		t.Errorf("list should be empty")
	}

	tdl.AppendToDo(&t1)
	if tdl.IsEmpty() {
		t.Errorf("list should not be empty")
	}

	if size := tdl.Size; size != 1 {
		t.Errorf("wrong count, expected 1 and got %d", size)
	}

	tdl.AppendToDo(&t2)
	tdl.AppendToDo(&t3)

	if size := tdl.Size; size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestGetById(t *testing.T) {
	tdl.clearList()

	tdl.AppendToDo(&t1)

	todo, err := tdl.GetToDoByID(id1)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	if id := todo.ID; id != id1 {
		t.Errorf("wrong entity returned, expected ID: 1 and got %d", id)
	}

	fail := uuid.New()
	todo, err = tdl.GetToDoByID(fail)
	if err == nil {
		t.Errorf("did not raise error properly for invalid GET")
	}
}

func TestGetByTitle(t *testing.T) {
	tdl.clearList()

	tdl.AppendToDo(&t1)

	todo, err := tdl.GetToDoByTitle(t1.Title)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	if title := todo.Title; title != t1.Title {
		t.Errorf("wrong entity returned, expected Title: %s and got %s", t1.Title, title)
	}

	fail := uuid.New()
	todo, err = tdl.GetToDoByID(fail)
	if err == nil {
		t.Errorf("did not raise error properly for invalid GET")
	}
}

func TestRemoveById(t *testing.T) {
	tdl.clearList()

	tdl.AppendToDo(&t1)
	err := tdl.RemoveToDoByID(id1)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if size := tdl.Size; size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}

	tdl.AppendToDo(&t1)
	tdl.AppendToDo(&t2)
	tdl.RemoveToDoByID(id1)
	if tdl.Head != &t2 {
		t.Errorf("head node was improperly removed")
	}

	tdl.AppendToDo(&t3)
	tdl.RemoveToDoByID(id3)
	if tdl.Tail != &t2 {
		t.Errorf("tail node was improperly removed")
	}

	tdl.clearList()
}

func TestUpdate(t *testing.T) {
	tdl.clearList()

	tdl.AppendToDo(&t1)
	expectedTime := time.Now()
	expectedToDo := ToDo{
		id1,
		"New Title",
		"New Description",
		expectedTime,
		time.Now(),
		true,
		nil,
		nil,
	}

	tdl.UpdateToDoEntity(expectedToDo)
	actualToDo, _ := tdl.GetToDoByID(id1)

	switch {
	case actualToDo.ID != expectedToDo.ID:
		t.Errorf("ID was updated")
	case actualToDo.Title != expectedToDo.Title:
		t.Errorf("Title was not updated")
	case actualToDo.Description != expectedToDo.Description:
		t.Errorf("Description was not updated")
	case actualToDo.StartedOn == expectedToDo.StartedOn:
		t.Errorf("Started on was updated")
	case actualToDo.CompletedOn == expectedToDo.CompletedOn:
		t.Errorf("Completed On was not updated")
	case !actualToDo.IsCompleted:
		t.Errorf("Completed not updated")
	}
}
