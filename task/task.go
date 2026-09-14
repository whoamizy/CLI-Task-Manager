package task

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

func NewTask(id int, title string) *Task {
	return &Task{
		ID:    id,
		Title: title,
		Done:  false,
	}
}

func (t *Task) Complete() {
	t.Done = true
}

func (t Task) String() string {
	checkMarker := "[ ]"
	if t.Done {
		checkMarker = "[x]"
	}

	return fmt.Sprintf("%s %d — %s", checkMarker, t.ID, t.Title)
}
