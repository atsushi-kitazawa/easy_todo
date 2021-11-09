package main

import (
	"fmt"

	"github.com/atsushi-kitazawa/easy_todo/domain"
	"github.com/atsushi-kitazawa/easy_todo/repository"
)

func main() {
    fmt.Println("easy todo")

    //task := entity.NewTask("aaa", "2021-11-09")
    //fmt.Println(task)

    // domain.InitTasks()

    repo := repository.NewInMemoryTaskRepository()
    t1, _ := domain.NewTask("aaa1", domain.Todo)
    repo.Save(*t1)
    t2, _ := domain.NewTask("aaa2", domain.Todo)
    repo.Save(*t2)
    t3, _ := domain.NewTask("bbb", domain.Progress)
    repo.Save(*t3)
    t4, _ := domain.NewTask("ccc", domain.Complete)
    repo.Save(*t4)
    repository.Print(*repo)

    task := repo.Find(1)
    fmt.Println(task)
    task = repo.Find(4)
    fmt.Println(task)
}
