package taskmanager

// Status represents the status of a task.
type Status int

const (
	StatusNotStarted           Status = iota // 0
	StatusRunning                           // 1
	StatusCompletedSuccessfully             // 2
	StatusError                             // 3
	StatusStopped                           // 4
)

// Task represents a task to be executed.
type Task struct {
	ID     int
	Action func() (interface{}, error)
	Status Status
	Result interface{}
	Err    error
}

func (task *Task) Execute(){
	task.Status = StatusRunning
	result,err := task.Action()
	if err != nil{
		task.Status = StatusError
		task.Err = err
		return
	}

	task.Status = StatusCompletedSuccessfully
	task.Result = result
}