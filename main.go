package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task-cli/task"
)

func main() {
	manager := task.NewManager()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите команду: ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Ошибка ввода:", err)
			}
			return
		}

		input := scanner.Text()
		parts := strings.Fields(input)
		if len(parts) == 0 {
			fmt.Println("Пустая команда, попробуйте другую")
			continue
		}
		cmd := parts[0]

		switch cmd {
		case "list":
			fmt.Println(manager.List())
		default:
			fmt.Println("Такой команды нет, попробуйте другую")
			continue
		}
	}
}
