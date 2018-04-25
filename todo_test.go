package main

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

var tdl ToDoList

var id1, _ = uuid.NewV4()
var id2, _ = uuid.NewV4()
var id3, _ = uuid.NewV4()
var id4, _ = uuid.NewV4()
var t1 = ToDo{id1, "First", "Testing", time.Now(), time.Time{}, false, nil, nil}
var t2 = ToDo{id2, "Second", "Testing", time.Now(), time.Time{}, false, nil, nil}
var t3 = ToDo{id3, "Third", "Testing", time.Now(), time.Time{}, false, nil, nil}
var t4 = ToDo{id4, "Fourth", "Testing", time.Now(), time.Time{}, false, nil, nil}

func TestIsEmpty(t *testing.T) {
	if tdl.Head == nil && !tdl.IsEmpty() {
		t.Errorf("list should be empty")
	}

	tdl.AppendToDo(&t1)

	if tdl.Head != nil && tdl.IsEmpty() {
		t.Errorf("list should NOT be empty")
	}

	tdl.clearList()
}

func TestAppend(t *testing.T) {
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
	todo, err := tdl.GetToDoByID(id1)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	if id := todo.ID; id != id1 {
		t.Errorf("wrong entity returned, expected ID: 1 and got %d", id)
	}

	fail, _ := uuid.NewV4()
	todo, err = tdl.GetToDoByID(fail)
	if err == nil {
		t.Errorf("did not raise error properly for invalid GET")
	}
}

func TestRemoveById(t *testing.T) {
	err := tdl.RemoveToDoByID(id1)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if size := tdl.Size; size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}

	if tdl.Head != &t2 {
		t.Errorf("head node was improperly removed")
	}

	tdl.RemoveToDoByID(id3)
	if tdl.Tail != &t2 {
		t.Errorf("tail node was improperly removed")
	}

	tdl.clearList()
}

func TestUpdate(t *testing.T) {
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
