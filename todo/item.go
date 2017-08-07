package todo

import (
	"encoding/json"
	"time"
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

func (item *Item) UnmarshalJSON(data []byte) error {
	base := struct {
		Id          uint      `json:"id"`
		Description string    `json:"description"`
		Status      Status    `json:"status"`
		Tags        []string  `json:"tags,omitempty"`
		Due         time.Time `json:"due,omitempty"`
	}{}

	err := json.Unmarshal(data, &base)

	if err != nil {
		return err
	}

	item.Id = base.Id
	item.Description = base.Description
	item.Status = base.Status

	if base.Tags == nil {
		item.Tags = []string{}
	} else {
		item.Tags = base.Tags
	}

	if base.Due.IsZero() {
		item.Due = time.Time{}
	} else {
		item.Due = base.Due
	}

	return nil
}

func (item Item) MarshalJSON() ([]byte, error) {
	var due string

	if item.Due.IsZero() {
		due = ""
	} else {
		due = item.Due.Format(time.ANSIC)
	}

	base := struct {
		Id          uint     `json:"id"`
		Description string   `json:"description"`
		Status      Status   `json:"status"`
		Tags        []string `json:"tags,omitempty"`
		Due         string   `json:"due,omitempty"`
	}{
		Id:          item.Id,
		Description: item.Description,
		Status:      item.Status,
		Tags:        item.Tags,
		Due:         due,
	}

	return json.Marshal(base)
}
