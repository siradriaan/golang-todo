package controllers

import (
	"github.com/siradriaan/go-todos/models/todos"
	"github.com/siradriaan/go-todos/views"
)

// FindAll the data
func FindAll() {
	data, err := todos.List()
	views.ViewAll(data, err)
}

func AddTodo(arg string) {
	data, err := todos.Add(arg)
	views.ViewSuccessNew(data, err)
}

func DelTodo(arg string) {
	data, err := todos.DeleteTodo(arg)
	views.ViewDel(data, err)
}

func DoneTodo(arg string) {
	data, err := todos.UpdateTodo(arg)
	views.ViewUpdate(data, err)
}
