package main

import (
	"fmt"

	"github.com/atsushi-kitazawa/easy_todo/entity"
)

func main() {
    fmt.Println("easy todo")

    //task := entity.NewTask("aaa", entity.Progress, "2021-11-09")
    //fmt.Println(task)

    entity.InitTasks()
}
