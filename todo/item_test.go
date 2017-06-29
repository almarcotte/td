package todo

import (
	"testing"
	"time"
)

func TestItem_IsPassedDue(t *testing.T) {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	tomorrow := today.AddDate(0, 0, 1)

	itemDue := Item{Due: yesterday}
	itemOnTime := Item{Due: tomorrow}
	itemNoDueDate := Item{}

	// Tests

	if !itemDue.IsPassedDue() {
		t.Fatal("Expecting the item to be passed due but isPassedDue() returned false.")
	}

	if itemOnTime.IsPassedDue() {
		t.Fatal("Expecting the item to not be passed due but isPassedDue() returned true.")
	}

	if itemNoDueDate.IsPassedDue() {
		t.Fatal("Expecting item with no due date to never be passed due but isPassedDue() returned true.")
	}
}

func TestItem_IsDone(t *testing.T) {
	itemDone := Item{Status: Completed}
	itemInProgress := Item{Status: InProgress}
	itemIncomplete := Item{Status: Incomplete}
	itemDefaultsIncomplete := Item{}

	if !itemDone.IsDone() {
		t.Fatal("Expecting completed item to be done.")
	}

	if itemInProgress.IsDone() {
		t.Fatal("Expecting item in progress to not be done.")
	}

	if itemIncomplete.IsDone() {
		t.Fatal("Expecting incomplete item to not be done.")
	}

	if itemDefaultsIncomplete.IsDone() {
		t.Fatal("Expecting items to default to incomplete.")
	}
}

func TestItem_HasTag(t *testing.T) {
	item := Item{Tags: []string{"important", "work", "golang"}}

	if item.HasTag("java") {
		t.Fatal("Expecting item to NOT have tag `java` but tag was found.")
	}

	if !item.HasTag("golang") {
		t.Fatal("Expecting item to have tag `golang` but not found.")
	}
}

func TestItem_AddTag(t *testing.T) {
	item := Item{Tags: []string{"golang"}}

	if item.HasTag("go") {
		t.Fatal("Expecting item to initially not have tag `go` but not found.")
	}

	item.AddTag("go")

	if !item.HasTag("go") {
		t.Fatal("Expecting item to have tag `go` but not found.")
	}

	item.AddTag("go") // Try adding the same tag again, make sure it's not in the list twice

	if len(item.Tags) != 2 {
		t.Fatal("Expecting 2 tags on item but found more.")
	}
}

func TestItem_RemoveTag(t *testing.T) {
	item := Item{Tags: []string{"golang", "java", "go", "work"}}

	if !item.HasTag("java") {
		t.Fatal("Expecting item to initially have tag `java` but not found.")
	}

	item.RemoveTag("java")

	if item.HasTag("java") {
		t.Fatal("Expecting item to no longer have tag `java` but found.")
	}
}
