package main

import (
	"testing"
)

var bstUUID Tree
var bstString Tree

func InsertByID(t *testing.T) {

}

func TestInsertByString(t *testing.T) {
	bstUUID.clearTree()
	bstString.clearTree()

	bstString.InsertByString(&t1)
	if bstString.Root == nil {
		t.Errorf("Insert By String failed: root is nil")
	}

	bstString.InsertByString(&t2)
	if bstString.Root.Left == nil && bstString.Root.Right == nil {
		t.Errorf("Insert By String failed: both leafs are nil")
	}
}

func TestFindByID(t *testing.T) {
	bstUUID.clearTree()
	bstString.clearTree()

	bstUUID.InsertByID(&t1)
	if node := bstUUID.FindByID(id1); node.ID != id1 {
		t.Errorf("Find By ID failed: wrong ID retrieved at root")
	}

	bstUUID.InsertByID(&t2)
	if node := bstUUID.FindByID(id2); node.ID != id2 {
		t.Errorf("Find By ID failed: wrong ID retrieved at leaf")
	}

	bstUUID.InsertByID(&t3)
	if node := bstUUID.FindByID(id3); node.ID != id3 {
		t.Errorf("Find By ID failed: wrong ID retrieved at leaf")
	}
}

func TestFindByString(t *testing.T) {
	bstUUID.clearTree()
	bstString.clearTree()

	bstString.InsertByString(&t1)
	if node := bstString.FindByString(t1.Title); node.Title != t1.Title {
		t.Errorf("Find By ID failed: wrong Title retrieved at root")
	}

	bstString.InsertByString(&t2)
	if node := bstString.FindByString(t2.Title); node.Title != t2.Title {
		t.Errorf("Find By ID failed: wrong Title retrieved at leaf")
	}

	bstString.InsertByString(&t3)
	if node := bstString.FindByString(t3.Title); node.Title != t3.Title {
		t.Errorf("Find By ID failed: wrong Title retrieved at leaf")
	}

}

func TestDeleteByID(t *testing.T) {
	bstUUID.clearTree()
	bstString.clearTree()

	bstUUID.InsertByID(&t1)
	bstUUID.DeleteByID(id1)
	if bstUUID.Root != nil {
		t.Errorf("Delete By ID failed: root node not deleted")
	}

	bstUUID.clearTree()

	bstUUID.InsertByID(&t1)
	bstUUID.InsertByID(&t2)
	bstUUID.DeleteByID(id1)
	if bstUUID.Root == nil {
		t.Errorf("Delete By String failed: root node not properly removed and replaced")
	}
}

func TestDeleteByString(t *testing.T) {
	bstUUID.clearTree()
	bstString.clearTree()

	bstUUID.InsertByString(&t1)
	bstUUID.DeleteByString(t1.Title)
	if bstUUID.Root != nil {
		t.Errorf("Delete By String failed: root node not deleted")
	}

	bstUUID.clearTree()

	bstUUID.InsertByString(&t1)
	bstUUID.InsertByString(&t2)
	bstUUID.DeleteByString(t1.Title)
	if bstUUID.Root == nil {
		t.Errorf("Delete By String failed: root node not properly removed and replaced")
	}
}
