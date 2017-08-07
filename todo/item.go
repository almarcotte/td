package todo

import (
	"time"
	"encoding/json"
)

type Status int

const (
	Incomplete Status = iota
	InProgress
	Completed
)

type Item struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	Tags        []string  `json:"tags"`
	Due         time.Time `json:"due,omitempty"`
}

// ChangeStatus sets the status of the item to the given status. Valid values are Incomplete, InProgress and Completed.
func (item *Item) ChangeStatus(status Status) {
	item.Status = status
}

// Complete changes the item's status to Complete. This is a shortcut for ChangeStatus(Complete)
func (item *Item) Complete() {
	item.Status = Completed
}

// InProgress changes the item's status to InProgress. This is a shortcut for ChangeStatus(InProgress)
func (item *Item) InProgress() {
	item.Status = InProgress
}

// Incomplete changes the item's status to Incomplete. This is a shortcut for ChangeStatus(Incomplete)
func (item *Item) Incomplete() {
	item.Status = Incomplete
}

// IsDone determines if the item is completed.
func (item *Item) IsDone() bool {
	return item.Status == Completed
}

// IsInProgress determines if the item is in progress.
func (item *Item) IsInProgress() bool {
	return item.Status == InProgress
}

// IsIncomplete determines if the item is not yet completed, meaning either in progress or incomplete.
func (item *Item) IsIncomplete() bool {
	return item.Status == Incomplete || item.Status == InProgress
}

// IsPassedDue determines if the item's due date is in the past. Items without a due date are never passed due.
func (item *Item) IsPassedDue() bool {
	return !item.Due.IsZero() && item.Due.Before(time.Now())
}

// AddTag adds the given tag to the item. If the tag is already applied we don't do anything.
func (item *Item) AddTag(tag string) {
	if tag != "" && !item.HasTag(tag) {
		item.Tags = append(item.Tags, tag)
	}
}

// HasTag returns true if the item has the given tag already applied, false otherwise.
func (item *Item) HasTag(tag string) bool {
	for i := range item.Tags {
		if item.Tags[i] == tag {
			return true
		}
	}

	return false
}

// RemoveTag removes the given tag from the item. If the tag wasn't applied it simply does nothing.
func (item *Item) RemoveTag(tag string) {
	tags := item.Tags[:0]

	for _, t := range item.Tags {
		if t != tag {
			tags = append(tags, t)
		}
	}

	item.Tags = tags
}

// ClearTags removes all the tags on an item
func (item *Item) ClearTags() {
	item.Tags = []string{}
}

// RemoveDueDate sets the Due to Time's zero value (time.Time{}) which is handled as an item without a due date
func (item *Item) RemoveDueDate() {
	item.Due = time.Time{}
}

// SetDueDate sets the item's due date / time
func (item *Item) SetDueDate(date time.Time) {
	item.Due = date
}

// SetDueTomorrow sets the due date to tomorrow. This is mostly a convenience function.
func (item *Item) SetDueTomorrow() {
	item.Due = time.Now().AddDate(0, 0, 1)
}

// SetDueNextWeek sets the due date to a week from today. This is mostly a convenience function.
func (item *Item) SetDueNextWeek() {
	item.Due = time.Now().AddDate(0, 0, 7)
}

// NewItemFromJson creates a new Item from the given json
func NewItemFromJson(data []byte) (*Item, error) {
	var item Item

	err := json.Unmarshal(data, &item)

	if err != nil {
		return nil, err
	}

	return &item, nil
}
