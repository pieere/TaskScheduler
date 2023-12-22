package main

import (
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	ID int
	Name string
	Time time.Time
	Function func() error
	Retry int
}

type Scheduler struct {
	tasks []Task
}

func Print() error {
	fmt.Println("看这里，打印啦")
	_, err := http.Get("fdsjfjfdldsjdsj")
	return err
}

func (s *Scheduler) Run(){
	for _, t := range s.tasks{
		go func(t Task){
			var err error
			for i := 0; i <= t.Retry; i++ {
				now := time.Now()
				duration := now.Sub(t.Time)
				if duration < 10*time.Second {
					err = t.Function()
					if err == nil{
					    break//完成任务，跳出循环
					} else {
						fmt.Printf("任务执行失败：%v, 进行第%d次尝试", err, i+1)
					}
				}
				time.Sleep(1*time.Second)



				}
			fmt.Print("warning:--------------任务执行失败----------------------------")
		}(t)
	}

}
func main() {
	var s Scheduler
	task1 := Task{ID:100000000, Name: "a",Time: time.Now(), Function: Print, Retry: 5}
	task2 := Task{ID:2, Name: "b",Time: time.Now().Add(5 * time.Second), Function: Print, Retry: 5}
	s.tasks = append(s.tasks, task1, task2)
	s.Run()
	time.Sleep(10 * time.Second)
}
