package domain

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewTask(t *testing.T) {
    taskA, _ := NewTask("aaa", Todo) 
    taskB, _ := NewTask("bbb", Progress) 
    taskC, _ := NewTask("ccc", Complete) 

    wantsA := Task{1, "aaa", Todo, ""}
    wantsB := Task{2, "bbb", Progress, ""}
    wantsC := Task{3, "ccc", Complete, ""}

    opt := cmp.AllowUnexported(Task{})
    fmt.Println(cmp.Equal(*taskA, wantsA, opt))
    fmt.Println(cmp.Equal(*taskB, wantsB, opt))
    fmt.Println(cmp.Equal(*taskC, wantsC, opt))
}

func TestModifyName(t *testing.T) {
    task, _ := NewTask("aaa", Todo)
    task.ModifyName("aaa change")

    wants := Task{4, "aaa change", Todo, ""}

    opt := cmp.AllowUnexported(Task{})
    fmt.Println(cmp.Equal(*task, wants, opt))
}

func TestModifyStatsu(t *testing.T) {
}

func TestDeleteTask(t *testing.T) {
    InitTasks()

    fmt.Println(alls)
    fmt.Println(todos)

    task := Task{5, "aaa", Todo, ""}
    task.DeleteTask()

    fmt.Println(alls)
    fmt.Println(todos)
}
