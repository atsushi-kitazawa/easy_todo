package entity

import (
	"fmt"
)

type Task struct {
    id int
    name string
    status Status
    completeDate string
}

type Status int

const (
    Todo Status = iota
    Progress
    Complete
)

type todos []Task
type progresses []Task
type completes []Task

var nextId = 0

func NewTask(name string, status Status, completeDate string) *Task {
    return &Task{id: getNextId(), name: name, status: status, completeDate: completeDate}
}

func InitTasks() {
    // [todo] from store file
    todoTasks := make(todos, 10)
    for i := 1; i <= 5; i++ {
	todoTasks = append(todoTasks, *NewTask("aaa", Todo, ""))
    }
    progressTasks := make(progresses, 10)
    for i := 6; i <= 10; i++ {
	progressTasks = append(progressTasks, *NewTask("bbb", Progress, ""))
    }
    completeTasks := make(completes, 10)
    for i := 11; i <= 15; i++ {
	completeTasks = append(completeTasks, *NewTask("ccc", Complete, ""))
    }

    fmt.Println(todoTasks)
    fmt.Println(progressTasks)
    fmt.Println(completeTasks)
}

func setNextId() {}

func getNextId() int {
    nextId++
    return nextId
}
