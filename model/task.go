package model

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task
var nextID = 1

func AddTask(t *Task) {
	t.ID = nextID
	nextID++
	tasks = append(tasks, *t)
}

func GetAllTasks() []Task {
	return tasks
}

func GetTaskByID(id int) (Task, bool) {
	for _, t := range tasks {
		if t.ID == id {
			return t, true
		}
	}
	return Task{}, false
}

func UpdateTask(id int, updated Task) bool {
	for i, t := range tasks {
		if t.ID == id {
			updated.ID = id
			tasks[i] = updated
			return true
		}
	}
	return false
}

func DeleteTask(id int) bool {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
