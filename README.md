# Task Manager

This repository provides a task manager for parallel task execution with implementation in Go. The task manager allows you to add tasks, execute them concurrently using multiple workers, and monitor their status and results.

## Installation

To use this task manager in your Go project, you can add it as a dependency using Go modules. Simply import the module in your code and import the necessary packages:

```go
import (
	"github.com/adivamshi/taskmanager"
)
```

Make sure you have Go 1.20 or a compatible version installed.

## Usage

### Task Manager Initialization

To use the task manager, you first need to initialize an instance. The task manager follows the singleton design pattern, so you can retrieve the instance using the `GetInstance()` method:

```go
tm := taskmanager.GetInstance()
```

### Adding Tasks

Tasks can be added to the task manager using the `AddTask()` method. A task is represented by the `Task` struct, which consists of an ID, an action function, and a status:

```go
task := &taskmanager.Task{
	ID:     1,
	Action: yourActionFunction,
}
tm.AddTask(task)
```

The action function should have the following signature: `func() (interface{}, error)`. It represents the task to be executed. You can define your own action function that performs the desired operations.

### Starting the Task Manager

Before the tasks can be executed, you need to start the task manager. This will define the number of workers that will be used to execute the tasks concurrently:

```go
tm.Start(numWorkers)
```

### Stopping the Task Manager

To stop the task manager and wait for all tasks to complete, you can use the `Stop()` method:

```go
tm.Stop()
```

### Waiting for Task Completion

If you want to wait for all tasks to complete without stopping the task manager, you can use the `WaitCompletion()` method:

```go
tm.WaitCompletion()
```

### Task Status and Results

Once the tasks have been executed, you can access their status and results. The `Task` struct has a `Status` field, which can have one of the following values:

- `StatusNotStarted`: The task has not started yet.
- `StatusRunning`: The task is currently running.
- `StatusCompletedSuccessfully`: The task has completed successfully.
- `StatusError`: The task encountered an error.
- `StatusStopped`: The task was stopped before completion.

You can access the status and result of a task after it has been executed:

```go
for _, completedTask := range tm.executedTasks {
	fmt.Println("Task ID:", completedTask.ID)
	fmt.Println("Status:", completedTask.Status)
	fmt.Println("Result:", completedTask.Result)
}
```

Note that the `executedTasks` field of the `TaskManager` struct stores all the executed tasks, including those that encountered errors or were stopped.

## Testing

The repository also includes test files (`taskmanager_test.go`) that demonstrate the usage and functionality of the task manager. You can run the tests using the `go test` command:

```bash
go test
```

The tests cover scenarios such as adding tasks, starting the task manager with multiple workers, and verifying task statuses and results.

## Contributing

Contributions to this repository are welcome. If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
