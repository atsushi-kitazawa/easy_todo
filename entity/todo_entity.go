package entity

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

func NewTask(id int, name string, status Status, completeDate string) *Task {
    return &Task{id: id, name: name, status: status, completeDate: completeDate}
}
