package domain

type Task struct {
	taskId          int
	title           string
	assignedMembers string
	status          string
	dueDate         string
}

func NewTask(title string) *Task {
	return &Task{
		title: title,
	}
}

func (t Task) Title() string {
	return t.title
}
