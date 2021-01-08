package views

import (
	"fmt"

	"github.com/siradriaan/go-todos/models/todos"
)

// ViewAll todos
func ViewAll(data []todos.Todo, err string) {
	if err != "" {
		fmt.Println(err)
	}

	var stats string
	for i := 0; i < len(data); i++ {
		if data[i].Status == false {
			stats = "[ ]"
		} else {
			stats = "[X]"
		}
		fmt.Println(data[i].ID, data[i].Name, stats)
	}
}

func ViewSuccessNew(data string, err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data)
}

func ViewDel(data string, err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data)
}

func ViewUpdate(data string, err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data)
}
