package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID int
	Name string
	Time time.Time
	Function func()
}

type Scheduler struct {
	tasks []Task
}

func Print() {
	fmt.Println("看这里，打印啦")
}

func (s *Scheduler) Run(){
	for _, t := range s.tasks{
		go func(t Task){
			now := time.Now()
			duration := now.Sub(t.Time)
			if duration < 10*time.Second {
				t.Function()
			}
		}(t)
	}

}
func main() {
	var s Scheduler
	task1 := Task{ID:100000000, Name: "a",Time: time.Now(), Function: Print}
	task2 := Task{ID:2, Name: "b",Time: time.Now().Add(5 * time.Second), Function: Print}
	s.tasks = append(s.tasks, task1, task2)
	s.Run()
	time.Sleep(15 * time.Second)
}
