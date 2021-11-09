package domain

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

func (t *Task) GetId() int {
    return t.id
}

func (t *Task) GetStatus() Status {
    return t.status
}

func (t *Task) Equal(another *Task) bool {
    return t.id == another.id
}

func setNextId(t []Task) {
    nextId = len(t)
}

func getNextId() int {
    nextId++ 
    return nextId
}
