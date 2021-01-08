package todos

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var todos = []Todo{}

// Todo model
type Todo struct {
	ID     int    `json:"id"`
	Name   string `json:"todo"`
	Status bool   `json:"status"`
}

// List all todos
func List() ([]Todo, string) {
	jsonFile, err := os.Open("../go-todos/data/todos.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)

	var todos []Todo

	json.Unmarshal(byteVal, &todos)
	if len(todos) == 0 {
		return nil, "you have no todo list"
	}
	return todos, ""
}

func Add(arg string) (string, error) {

	if arg == "" {
		return "put something will ya", nil
	}

	todos, _ := List()

	todos = append(todos, Todo{
		ID:     len(todos) + 1,
		Name:   arg,
		Status: false,
	})

	b, _ := json.Marshal(todos)

	err := ioutil.WriteFile("../go-todos/data/todos.json", b, os.ModePerm)

	if err != nil {
		return "", err
	}

	return "success adding new todo", nil
}

func DeleteTodo(id string) (string, error) {

	todos, _ := List()

	in, _ := strconv.Atoi(id)

	if in <= 0 || in > len(todos) {
		return "", errors.New("cannot found todo")
	}

	copy(todos[in-1:], todos[in:])

	todos[len(todos)-1] = Todo{}

	todos = todos[:len(todos)-1]

	for i := 0; i < len(todos); i++ {
		todos[i].ID = i + 1
	}

	t, _ := json.Marshal(todos)

	err := ioutil.WriteFile("../go-todos/data/todos.json", t, os.ModePerm)
	if err != nil {
		return "", err
	}

	return "delete succeed", nil
}

func UpdateTodo(id string) (string, error) {
	todos, _ := List()

	in, _ := strconv.Atoi(id)

	if in <= 0 || in > len(todos) {
		return "", errors.New("not found")
	}

	todos[in-1].Status = true

	t, _ := json.Marshal(todos)

	err := ioutil.WriteFile("../go-todos/data/todos.json", t, os.ModePerm)

	if err != nil {
		return "", err
	}

	return "update succeed", nil

}
