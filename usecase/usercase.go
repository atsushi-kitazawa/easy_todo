package usecase

import "github.com/atsushi-kitazawa/easy_todo/domain"

func AddTask(name string, status domain.Status) {
    _, err := domain.NewTask(name, status)
    if err != nil {
	panic(err)
    }

}
