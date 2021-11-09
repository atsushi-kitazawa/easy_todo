package main

import (
	"fmt"

	"github.com/atsushi-kitazawa/easy_todo/entity"
)

func main() {
    fmt.Println("easy todo")

    task := entity.NewTask(1, "aaa", entity.Progress, "2021-11-09")
    fmt.Println(task)
}
