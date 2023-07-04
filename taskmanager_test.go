package taskmanager

import (
	"testing"
	"time"
)

var taskFunc = func(name string, duration time.Duration) func() (interface{}, error) {
		return func() (interface{}, error) {
			time.Sleep(duration)
			return name, nil
		}
	}

func TestAddTask(t *testing.T) {
	tm := GetInstance()
	tm.Start(100)
	task := &Task{
		ID: 1,
		Action: taskFunc("Task 1", 5*time.Second),
	}
	tm.AddTask(task)
}

func TestTaskManager(t *testing.T){
	tm := GetInstance()
	tm.Start(1)

	task := &Task{
		ID: 1,
		Action: taskFunc("Task 1", 5*time.Millisecond),
	}
	tm.AddTask(task)
	tm.WaitCompletion()
	for _, completedTask := range tm.executedTasks{
		if completedTask.Status != StatusCompletedSuccessfully {
			t.Errorf("Task 1 status: expected %d, got %d", StatusCompletedSuccessfully, completedTask.Status)
		}
		if completedTask.Result != "Task 1" {
			t.Errorf("Task 1 result: expected %s, got %v", "Task 1", completedTask.Result)
		}
	}
}