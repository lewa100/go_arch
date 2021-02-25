package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.New("Todolist").ParseFiles("template.html"))

func TodoListHandler(w http.ResponseWriter, r *http.Request) {
	todo := GetTodoList()
	if err := tmpl.Execute(w, todo); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func SwitchDoneHandler(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Query().Get("task_id")
	if taskID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todoList := GetTodoList()
	for _, task := range todoList.Tasks {
		if strconv.Itoa(task.ID) == taskID {
			task.IsDone = !task.IsDone
		}
	}

	todo := GetTodoList()
	if err := tmpl.Execute(w, todo); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	serv := http.NewServeMux()

	serv.HandleFunc("/", TodoListHandler)
	serv.HandleFunc("/switch/", SwitchDoneHandler)
	log.Println("Starting server on port 8055...")
	http.ListenAndServe(":8055", serv)
}
