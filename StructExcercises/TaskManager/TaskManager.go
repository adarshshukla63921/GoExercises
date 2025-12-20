package main

import "fmt"

type Task struct {
	ID       int
	Name     string
	Priority int
	Done     bool
}

func MarkDone(t *Task) {
	if t == nil {
		return
	}
	if(t.Done){
		return
	}
	t.Done = true
}

func UpdatePriority(t *Task, newPriority int){
	if(t == nil){
		return
	}
	if(newPriority < 1 || newPriority > 5){
		return
	}
	t.Priority = newPriority
}

// Composite Function must 'guard' before calling other functions, even if other functions guard.
func EscalateTask(t *Task){
	// we need to derefernce here, since we are using t.Priority.
	// we would also do that if we use *t.
	if( t == nil){
		return
	}
	MarkDone(t)

	if t.Priority < 5 {
		UpdatePriority(t, t.Priority+1)
	}
}
func main() {
	task := Task{101, "Practice Go", 5, false}
	MarkDone(&task)
	fmt.Println("Using Markdone func : ",task)

	task2 := Task{102, "Read Book", 3, false}
	UpdatePriority(&task2, 4)
	fmt.Println("Using UpdatePriority func : ",task2)

	task3 := Task{103, "Write Code", 2, false}
	EscalateTask(&task3)
	fmt.Println("Using EsclateTask func : ",task3)

	task4 := Task{104, "Impossible Task", 5 , true}
	EscalateTask(&task4)
}