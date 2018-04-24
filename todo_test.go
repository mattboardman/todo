package main

import (
	"testing"
	"time"

	"github.com/nu7hatch/gouuid"
)

var tdl ToDoList

var id1, _ = uuid.NewV4()
var id2, _ = uuid.NewV4()
var id3, _ = uuid.NewV4()
var id4, _ = uuid.NewV4()
var t1 = ToDo{*id1, 0, "First", "Testing", time.Now(), time.Time{}, nil, nil}
var t2 = ToDo{*id2, 0, "Second", "Testing", time.Now(), time.Time{}, nil, nil}
var t3 = ToDo{*id3, 0, "Third", "Testing", time.Now(), time.Time{}, nil, nil}
var t4 = ToDo{*id4, 2, "Fourth", "Testing", time.Now(), time.Time{}, nil, nil}

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
	todo, err := tdl.GetToDoByID(*id1)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	if id := todo.ID; id != *id1 {
		t.Errorf("wrong entity returned, expected ID: 1 and got %d", id)
	}

	fail, _ := uuid.NewV4()
	todo, err = tdl.GetToDoByID(*fail)
	if err == nil {
		t.Errorf("did not raise error properly for invalid GET")
	}
}

func TestInsertToDo(t *testing.T) {
	tdl.InsertToDo(&t4)
	if tdl.Head.Next != &t4 {
		t.Errorf("Insertion failed")
	}

	if order := t4.Next.Order; order != 3 {
		t.Errorf("Reorder failed, expected 3 got %d", order)
	}
}

func TestRemoveById(t *testing.T) {
	err := tdl.RemoveToDoByID(*id1)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if size := tdl.Size; size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}

	if tdl.Head != &t2 {
		t.Errorf("head node was improperly removed")
	}

	tdl.RemoveToDoByID(*id3)
	if tdl.Tail != &t2 {
		t.Errorf("tail node was improperly removed")
	}

	if t2.Order != 1 {
		t.Errorf("reorder failed")
	}
}
