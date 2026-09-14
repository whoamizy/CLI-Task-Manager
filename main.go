package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task-cli/task"
)

var COMMANDS = map[string]string{
	"add":    "Добавить новую задачу",
	"list":   "Показать список всех задач",
	"done":   "Отметить задачу как выполненную",
	"delete": "Удалить задачу",
	"find":   "Найти задачи по тексту",
	"help":   "Показать список доступных команд",
	"exit":   "Завершить программу",
}

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
		case "add":
			userTask := parts[1:]
			if len(userTask) == 0 {
				fmt.Println("Такое нельзя добавить")
				continue
			}
			task := manager.Add(strings.Join(userTask, " "))

			fmt.Println("Добавлена задача:")
			fmt.Println(task)
		case "help":
			for key, value := range COMMANDS {
				fmt.Println(key, "-", value)
			}
		case "list":
			tasks := manager.List()
			if len(tasks) == 0 {
				fmt.Println("Список пуст")
				continue
			}
			for _, task := range tasks {
				fmt.Println(task)
			}
		case "exit":
			fmt.Println("До скорого!")
			return
		default:
			fmt.Println("Такой команды нет, попробуйте другую")
			continue
		}
	}
}
