package main

import (
	"errors"
)

type Node struct {
	ToDo  *ToDo
	Left  *Node
	Right *Node
}

func (n *Node) Insert(todo *ToDo) error {
	if n == nil {
		return errors.New("Tree has not been instantiated")
	}

	switch {
	case todo.ID.ID() == n.ToDo.ID.ID():
		return nil
	case todo.ID.ID() < n.ToDo.ID.ID():
		if n.Left == nil {
			n.Left = &Node{ToDo: todo}
			return nil
		}
		return n.Left.Insert(todo)
	case todo.ID.ID() > n.ToDo.ID.ID():
		if n.Right == nil {
			n.Right = &Node{ToDo: todo}
			return nil
		}
		return n.Right.Insert(todo)
	}
	return nil
}
