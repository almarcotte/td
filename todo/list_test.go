package todo

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

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

func TestNewListFromJson(t *testing.T) {
	data := []byte(`{
  "lastid": 3,
  "items": [
    {"id": 1, "description": "Test Item #1", "status": 0, "tags": ["test", "golang"]},
    {"id": 2, "description": "Test Item #2", "status": 0, "tags": ["test", "golang", "test2"]},
    {"id": 3, "description": "Test Item #2", "status": 0}
  ]
}`)
	var list List

	err := json.Unmarshal(data, &list)

	if err != nil {
		t.Fatalf("Unexpected errors while unmarshalling json: %s", err.Error())
	}

	if list.Count() != 3 {
		t.Fatalf("Expected 3 items, got %d", list.Count())
	}

	if list.LastId != 3 {
		t.Fatalf("Expected last used id to be 3, got %d", list.LastId)
	}

	if item := list.Get(1); item.Description != "Test Item #1" {
		t.Fatalf("Expected item 1 to have description `Test Item #1`, got %s", item.Description)
	}

	if item := list.Get(2); !item.HasTag("test2") {
		t.Fatalf("Expected item 2 to have tag `test2` but not found, got %v+", item.Tags)
	}
}

func TestList_ToJson(t *testing.T) {
	list := NewList()
	item := Item{Description: "Item #1", Tags: []string{"golang", "is", "awesome"}}

	item2 := Item{Description: "Item #2"}
	item2.SetDueNextWeek()

	dueDate := item2.Due.Format(time.ANSIC)

	list.Add(item)
	list.Add(item2)

	content := `{"items":[{"id":1,"description":"Item #1","status":0,"tags":["golang","is","awesome"]},{"id":2,"description":"Item #2","status":0,"due":"%s"}],"lastid":2}`
	expected := string([]byte(fmt.Sprintf(content, dueDate)))
	output, err := json.Marshal(list)

	if err != nil {
		t.Fatalf("Unexpected error from marshalling: %s", err.Error())
	}

	if string(output) != expected {
		t.Fatalf("Unexpected result JSON.\nExpected: %s\nGot     : %s", expected, string(output))
	}
}
