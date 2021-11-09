package domain

type allTasks []Task
type todoTasks []Task
type progressTasks []Task
type completeTasks []Task

var alls = make(allTasks,0)
var todos = make(todoTasks,0)
var progresses = make(progressTasks,0)
var completes = make(completeTasks,0)

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
