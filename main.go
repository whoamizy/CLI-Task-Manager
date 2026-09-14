package main

import (
	"fmt"
	"task-cli/task"
)

func main() {
	task := task.NewTask(1, "Learn Go")

	fmt.Println(task)

	task.Complete()

	fmt.Println(task)
}
