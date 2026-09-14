package task

import (
	"errors"
	"strings"
)

type Manager struct {
	tasks  []*Task
	nextID int
}

func NewManager() *Manager {
	return &Manager{
		tasks:  make([]*Task, 0),
		nextID: 1,
	}
}

func (m *Manager) Add(title string) *Task {
	newTask := NewTask(m.nextID, title)
	m.tasks = append(m.tasks, newTask)
	m.nextID++
	return newTask
}

func (m *Manager) List() []*Task {
	return m.tasks
}

func (m *Manager) Complete(id int) error {
	for _, task := range m.tasks {
		if task.ID == id {
			task.Complete()
			return nil
		}
	}

	return errors.New("task not found")
}

func (m *Manager) Delete(id int) error {
	for index, task := range m.tasks {
		if task.ID == id {
			m.tasks = append(m.tasks[:index], m.tasks[index+1:]...)
			return nil
		}
	}

	return errors.New("task not found")
}

func (m *Manager) Find(query string) []*Task {
	res := []*Task{}
	query = strings.TrimSpace(strings.ToLower(query))
	if query == "" {
		return res
	}

	for _, task := range m.tasks {
		if strings.Contains(strings.ToLower(task.Title), query) {
			res = append(res, task)
		}
	}

	return res
}
