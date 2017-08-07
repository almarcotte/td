package todo

import "encoding/json"

type List struct {
	Items    []Item `json:"items"`
	IsGlobal bool   `json:"-"` // Omit from json
	LastId   uint   `json:"lastid"`
}

func NewList() *List {
	return &List{LastId: 0}
}

func (l *List) IsEmpty() bool {
	return len(l.Items) == 0
}

func (l *List) Count() int {
	return len(l.Items)
}

func (l *List) Add(item Item) {
	l.LastId++
	item.Id = l.LastId

	l.Items = append(l.Items, item)
}

// Get returns the item from this list with the matching id, nil if not found
func (l *List) Get(id uint) *Item {
	if id > l.LastId {
		return nil
	}

	for _, item := range l.Items {
		if item.Id == id {
			return &item
		}
	}

	return nil
}

// NewListFromJson returns a list from the given json
func NewListFromJson(data []byte) (*List, error) {
	var list List

	err := json.Unmarshal(data, &list)

	if err != nil {
		return nil, err
	}

	return &list, nil
}
