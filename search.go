package main

import (
	"errors"
	"strings"
	"time"

	uuid "github.com/google/uuid"
)

// Node struct contains:
// 1. Reference to a ToDo item
// 2. Reference to Left Node
// 3. Reference to Right Node
type Node struct {
	ToDo  *ToDo
	Left  *Node
	Right *Node
}

// InsertByID inserts a ToDo item by its UUID
func (n *Node) InsertByID(todo *ToDo) error {
	if n == nil {
		return errors.New("Cannot INSERT: Tree has not been instantiated")
	}

	switch {
	case todo.ID.ID() == n.ToDo.ID.ID():
		return nil
	case todo.ID.ID() < n.ToDo.ID.ID():
		if n.Left == nil {
			n.Left = &Node{ToDo: todo}
			return nil
		}
		return n.Left.InsertByID(todo)
	case todo.ID.ID() > n.ToDo.ID.ID():
		if n.Right == nil {
			n.Right = &Node{ToDo: todo}
			return nil
		}
		return n.Right.InsertByID(todo)
	}
	return nil
}

// InsertByString inserts a ToDo item by a string value
// If the Title exists then the Title is inserted as another
// tree by ID
func (n *Node) InsertByString(todo *ToDo) error {
	if n == nil {
		return errors.New("Cannot INSERT: Tree has not been instantiated")
	}

	switch {
	case strings.ToLower(todo.Title) == strings.ToLower(n.ToDo.Title):
		if todo.ID.ID() != n.ToDo.ID.ID() {
			n.InsertByID(todo)
		}
		return nil
	case strings.ToLower(todo.Title) < strings.ToLower(n.ToDo.Title):
		if n.Left == nil {
			n.Left = &Node{ToDo: todo}
			return nil
		}
		return n.Left.InsertByString(todo)
	case strings.ToLower(todo.Title) > strings.ToLower(n.ToDo.Title):
		if n.Right == nil {
			n.Right = &Node{ToDo: todo}
			return nil
		}
		return n.Right.InsertByString(todo)
	}
	return nil
}

// FindByID method traverses the Tree recursively
// until it finds the matching UUID
func (n *Node) FindByID(id uuid.UUID) *ToDo {
	if n == nil {
		return nil
	}

	switch {
	case id.ID() == n.ToDo.ID.ID():
		return n.ToDo
	case id.ID() < n.ToDo.ID.ID():
		return n.Left.FindByID(id)
	default:
		return n.Right.FindByID(id)
	}
}

// FindByString method traverses the Tree recursively
// until it finds the matching string value
func (n *Node) FindByString(value string) *ToDo {
	if n == nil {
		return nil
	}

	time.Sleep(10 * time.Millisecond)
	switch {
	case strings.ToLower(value) == strings.ToLower(n.ToDo.Title):
		return n.ToDo
	case strings.ToLower(value) < strings.ToLower(n.ToDo.Title):
		return n.Left.FindByString(value)
	default:
		return n.Right.FindByString(value)
	}
}

// findMax finds the right most element of the Tree
func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

// replaceNode replaces the node with its right most value in the subtree
func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("Cannot REPLACE: node to replace is nil")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

// DeleteByID deletes a Node by its UUID
func (n *Node) DeleteByID(id uuid.UUID, parent *Node) error {
	if n == nil {
		return errors.New("Cannot DELETE: ID does not exist in the tree")
	}

	switch {
	case id.ID() < n.ToDo.ID.ID():
		return n.Left.DeleteByID(id, n)
	case id.ID() > n.ToDo.ID.ID():
		return n.Right.DeleteByID(id, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}
		replacement, replParent := n.Left.findMax(n)
		n.ToDo = replacement.ToDo
		return replacement.DeleteByID(replacement.ToDo.ID, replParent)
	}
}

// DeleteByString deletes a Node by a string value
func (n *Node) DeleteByString(value string, parent *Node) error {
	if n == nil {
		return errors.New("Cannot DELETE: ID does not exist in the tree")
	}

	switch {
	case strings.ToLower(value) < strings.ToLower(n.ToDo.Title):
		return n.Left.DeleteByString(value, n)
	case strings.ToLower(value) > strings.ToLower(n.ToDo.Title):
		return n.Right.DeleteByString(value, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}
		replacement, replParent := n.Left.findMax(n)
		n.ToDo = replacement.ToDo
		return replacement.DeleteByString(replacement.ToDo.Title, replParent)
	}
}

// Tree struct contains the root node
type Tree struct {
	Root *Node
}

// InsertByID inserts a ToDo item by its UUID
func (t *Tree) InsertByID(todo *ToDo) error {
	if t.Root == nil {
		t.Root = &Node{ToDo: todo}
		return nil
	}
	return t.Root.InsertByID(todo)
}

// InsertByString inserts a ToDo item by its string value
func (t *Tree) InsertByString(todo *ToDo) error {
	if t.Root == nil {
		t.Root = &Node{ToDo: todo}
		return nil
	}
	return t.Root.InsertByString(todo)
}

// FindByID finds a Node by its UUID
func (t *Tree) FindByID(id uuid.UUID) *ToDo {
	if t.Root == nil {
		return nil
	}
	return t.Root.FindByID(id)
}

// FindByString finds a Node by its string value
func (t *Tree) FindByString(value string) *ToDo {
	if t.Root == nil {
		return nil
	}
	return t.Root.FindByString(value)
}

// DeleteByID removes a Node by its UUID
func (t *Tree) DeleteByID(id uuid.UUID) error {
	if t.Root == nil {
		return errors.New("Cannot DELETE: tree is empty")
	}
	fakeParent := &Node{Right: t.Root}
	err := t.Root.DeleteByID(id, fakeParent)
	if err != nil {
		return err
	}

	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}

// DeleteByString removes a Node by its string value
func (t *Tree) DeleteByString(value string) error {
	if t.Root == nil {
		return errors.New("Cannot DELETE: tree is empty")
	}
	fakeParent := &Node{Right: t.Root}
	err := t.Root.DeleteByString(value, fakeParent)
	if err != nil {
		return err
	}

	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}

func (t *Tree) clearTree() {
	t.Root = nil
}
