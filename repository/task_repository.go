package repository

import (
	"fmt"

	"github.com/atsushi-kitazawa/easy_todo/domain"
)

type ITaskRespository interface {
	Find(id int)
	Save(domain.Task)
	Update(domain.Task)
	Delete(domain.Task)
}

type InMemoryTaskRepository struct {
	alls       []domain.Task
	todos      []domain.Task
	progresses []domain.Task
	completes  []domain.Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		alls:       make([]domain.Task, 0),
		todos:      make([]domain.Task, 0),
		progresses: make([]domain.Task, 0),
		completes:  make([]domain.Task, 0),
	}
}

func (repo *InMemoryTaskRepository) Find(id int) *domain.Task {
	for _, v := range repo.alls {
		if v.GetId() == id {
			return &v
		}
	}
	return nil
}

func (repo *InMemoryTaskRepository) Save(t domain.Task) {
	repo.alls = append(repo.alls, t)

	switch t.GetStatus() {
	case domain.Todo:
		repo.todos = append(repo.todos, t)
	case domain.Progress:
		repo.progresses = append(repo.progresses, t)
	case domain.Complete:
		repo.completes = append(repo.completes, t)
	}
}

func (repo *InMemoryTaskRepository) Update(t domain.Task) {
	for i, v := range repo.alls {
		if v.Equal(&t) {
			repo.alls[i] = t
			break
		}
	}
}

func (repo *InMemoryTaskRepository) Delete(t domain.Task) {
	for i, v := range repo.alls {
		if v.GetId() == t.GetId() {
			removeElement(&repo.alls, i)
			break
		}
	}

	switch t.GetStatus() {
	case domain.Todo:
		for i, v := range repo.todos {
			if v.GetId() == t.GetId() {
				removeElement(&repo.todos, i)
				break
			}
		}
	case domain.Progress:
		for i, v := range repo.progresses {
			if v.GetId() == t.GetId() {
				removeElement(&repo.progresses, i)
				break
			}
		}
	case domain.Complete:
		for i, v := range repo.completes {
			if v.GetId() == t.GetId() {
				removeElement(&repo.completes, i)
				break
			}
		}
	}
}

func removeElement(slice *[]domain.Task, index int) {
	(*slice)[index] = (*slice)[len((*slice))-1]
	(*slice) = (*slice)[:len((*slice))-1]
}

func Print(repo InMemoryTaskRepository) {
    fmt.Println(repo.alls)
    fmt.Println(repo.todos)
    fmt.Println(repo.progresses)
    fmt.Println(repo.completes)
}

func (repo InMemoryTaskRepository) Load() {
	//for i := 1; i <= 5; i++ {
	//	t, _ := domain.NewTask("aaa", domain.Todo)
	//	repo.todos = append(repo.todos, *t)
	//}
	//for i := 6; i <= 10; i++ {
	//	t, _ := domain.NewTask("bbb", domain.Progress)
	//	repo.progresses = append(repo.progresses, *t)
	//}
	//for i := 11; i <= 15; i++ {
	//	t, _ := domain.NewTask("ccc", domain.Complete)
	//	repo.completes = append(repo.completes, *t)
	//}
	//for _, v := range repo.todos {
	//	repo.alls = append(repo.alls, v)
	//}
	//for _, v := range repo.progresses {
	//	repo.alls = append(repo.alls, v)
	//}
	//for _, v := range repo.completes {
	//	repo.alls = append(repo.alls, v)
	//}
	//fmt.Println(repo.todos)
	//fmt.Println(repo.progresses)
	//fmt.Println(repo.completes)
	//fmt.Println(repo.alls)
}
