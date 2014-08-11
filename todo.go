package main

import (
	"github.com/nu7hatch/gouuid"
)

func newId() string {
	id, _ := uuid.NewV4()
	return id.String()
}

type TodoItem struct {
	Id        string `json:"-"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	Text      string `json:"text"`
}

func (i *TodoItem) Update(item TodoItem) *TodoItem {
	i.Title = item.Title
	i.Completed = item.Completed
	i.Order = item.Order
	i.Text = item.Text
	return i
}

type Todo map[string]*TodoItem

func (t Todo) All() []*TodoItem {
	items := []*TodoItem{}
	for _, item := range t {
		items = append(items, item)
	}
	return items
}

func (t Todo) Find(id string) *TodoItem {
	for _, item := range t {
		if item.Id == id {
			return item
		}
	}

	return nil
}

func (t Todo) Create(item TodoItem, fqdn func(string) string) *TodoItem {
	item.Id = newId()
	item.Url = fqdn("/todos/" + item.Id)
	t[item.Id] = &item
	return &item
}

func (t Todo) Update(id string, updatedItem TodoItem) *TodoItem {
	if item := t.Find(id); item != nil {
		return item.Update(updatedItem)
	} else {
		return nil
	}
}

func (t Todo) DeleteAll() string {
	for k := range t {
		delete(t, k)
	}
	return ""
}

func (t Todo) Delete(id string) string {
	for k := range t {
		if k == id {
			delete(t, k)
		}
	}
	return ""
}
