package main

type TodoList struct {
	Title string
	Tasks []*Task
}

type Task struct {
	ID     int
	Name   string
	IsDone bool
}

var List *TodoList

func GetTodoList() *TodoList {
	if List == nil {
		List = &TodoList{
			Title: "Мой таск трекер",
			Tasks: []*Task{
				{
					ID:     1,
					Name:   "Убраться",
					IsDone: false,
				},
				{
					ID:     2,
					Name:   "Поработать",
					IsDone: true,
				},
				{
					ID:     3,
					Name:   "Выспаться",
					IsDone: false,
				},
			},
		}
	}
	return List
}
