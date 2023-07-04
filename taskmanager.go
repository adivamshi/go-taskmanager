package taskmanager

import (
	"sync"
	"sync/atomic"
)

type TaskManager struct{
	tasks chan *Task
	executedTasks []*Task
	workers int
	status Status
}

var once sync.Once
var mutex sync.Mutex
var instance *TaskManager
var workerGroup sync.WaitGroup
var runningWorkers int32

func GetInstance() *TaskManager{
	once.Do(func() {
		instance = &TaskManager{
			tasks: make(chan *Task),
			workers: 1,
			status: StatusNotStarted,
		}
	})
	return instance
}

func (tm *TaskManager) AddTask(task *Task){
	if tm.status == StatusRunning{
		tm.Start(1)
	}
	task.Status = StatusNotStarted
	tm.tasks <- task
}

func (tm *TaskManager) Start(workers int){
	tm.workers = workers

	for i := 0; i < workers; i++ {
		workerGroup.Add(1)
		atomic.AddInt32(&runningWorkers, 1)
		go tm.runWorker()
	}
}

func (tm *TaskManager) runWorker(){
	defer workerGroup.Done()

	for{
		select {
		case task := <-tm.tasks:
			task.Execute()
			tm.executedTasks = append(tm.executedTasks, task)
		default:
			mutex.Lock()
			if tm.workers < int(runningWorkers){
				atomic.AddInt32(&runningWorkers,-1)
				return
			}
			mutex.Unlock()
			if tm.status == StatusStopped{
				return
			}
		}
	}
}

func (tm *TaskManager) Stop(){
	tm.status=StatusStopped
	close(tm.tasks)
	tm.WaitCompletion()
	tm.workers=0
}

func (tm *TaskManager) WaitCompletion(){
	workerGroup.Wait()
	tm.workers = 0
	tm.status = StatusCompletedSuccessfully
}
