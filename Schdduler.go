package main

import (
	"fmt"
	"sort"
	"time"
)

type Task struct {
	ID int
	Name string
	Time time.Time
	Function func()
	Priority int
}

type Scheduler struct {
	tasks []Task
}

func Print() {
	fmt.Println("看这里，打印啦")
}

func (s *Scheduler) Run(){
	sort.Slice(s.tasks, func(i, j int) bool {
		return s.tasks[i].Priority < s.tasks[j].Priority
	})
	fmt.Print(s.tasks)
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
	task1 := Task{ID:1, Name: "a",Time: time.Now(), Function: Print, Priority: 8000}
	task2 := Task{ID:2, Name: "b",Time: time.Now().Add(5 * time.Second), Function: Print, Priority: 100}
	s.tasks = append(s.tasks, task1, task2)
	s.Run()
	time.Sleep(15 * time.Second)
}
