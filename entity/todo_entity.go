package entity

import (
	"errors"
	_ "fmt"
	"time"
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

var alls = make([]Task,0)
var todos = make([]Task,0)
var progresses = make([]Task,0)
var completes = make([]Task,0)

var nextId = 0

func NewTask(name string, status Status) (*Task, error) {
    if len(name) == 0 {
	return nil, errors.New("Please specify name with at least one character")	
    }

    return &Task{id: getNextId(), name: name, status: status, completeDate: ""}, nil
}

func (t *Task) ModifyName(name string) {
    if len(name) == 0 {
	return
    }

    t.name = name
}

func (t *Task) ModifyStatus(status Status) {
    t.status = status

    if status == Complete {
	t.completeDate = time.Now().String()
    }
}

func (t *Task) DeleteTask() {
    switch t.status {
    case Todo:
	for i, v := range todos {
	    if v.id == t.id {
		removeElement(&todos, i)
		removeElement(&alls, i)
	    }
	}
    case Progress:
	for i, v := range progresses {
	    if v.id == t.id {
		removeElement(&progresses, i)
		removeElement(&alls, i)
	    }
	}
    case Complete:
	for i, v := range completes {
	    if v.id == t.id {
		removeElement(&completes, i)
		removeElement(&alls, i)
	    }
	}
    }
}

func setNextId(t []Task) {
    nextId = len(t)
}

func getNextId() int {
    nextId++ 
    return nextId
}

func removeElement(slice *[]Task, index int) {
    (*slice)[index] = (*slice)[len((*slice))-1]
    (*slice) = (*slice)[:len((*slice))-1]
}

func InitTasks() {
    // [todo] from store file
    for i := 1; i <= 5; i++ {
	t, _  := NewTask("aaa", Todo)
	todos = append(todos, *t)
    }
    for i := 6; i <= 10; i++ {
	t, _ := NewTask("bbb", Progress)
	progresses = append(progresses, *t)
    }
    for i := 11; i <= 15; i++ {
	t, _ := NewTask("ccc", Complete)
	completes = append(completes, *t)
    }
    for _, v := range todos {
	alls = append(alls, v)
    }
    for _, v := range progresses {
	alls = append(alls, v)
    }
    for _, v := range completes {
	alls = append(alls, v)
    }
    //fmt.Println(todos)
    //fmt.Println(progresses)
    //fmt.Println(completes)
    //fmt.Println(alls)
}
