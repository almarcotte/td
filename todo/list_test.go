package todo

import "testing"

func TestList_IsEmpty(t *testing.T) {
	list := NewList()

	if !list.IsEmpty() {
		t.Fatalf("Expected list to be empty")
	}
}

func TestList_Add(t *testing.T) {
	list := NewList()
	item := Item{Description: "Testing", Status: Incomplete}

	if !list.IsEmpty() {
		t.Fatalf("Expected new list to be empty")
	}

	list.Add(item)

	if list.IsEmpty() {
		t.Fatalf("Expected list to have one item but it is empty")
	}
}

func TestList_Get(t *testing.T) {
	list := NewList()
	item := Item{Description: "Item #1"}
	item2 := Item{Description: "Item #2"}

	list.Add(item)
	list.Add(item2)

	if fetched := list.Get(2); fetched.Id != 2 {
		t.Fatalf("Expected the returned item's ID to be 2, got %s", fetched.Id)
	}

	if fetched := list.Get(2); fetched.Description != "Item #2" {
		t.Fatalf("Expected the returned item's description to be `Item #2`, got %s", fetched.Description)
	}
}
