package main

import (
	"fmt"

	"github.com/atsushi-kitazawa/easy_todo/domain"
)

func main() {
    fmt.Println("easy todo")

    //task := entity.NewTask("aaa", "2021-11-09")
    //fmt.Println(task)

    domain.InitTasks()
}
