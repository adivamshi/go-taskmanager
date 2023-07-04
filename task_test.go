package taskmanager

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	task := Task{
		ID: 1,
		Action: func() (interface{}, error) {
			fmt.Println("This is inside action")
			return  true,nil
		},
	}

	task.Execute()
	if task.Result != true{
		t.Error("Error executing action")
	}
	if task.Err != nil{
		t.Errorf("Error executing action:%v",task.Err)
	}
}