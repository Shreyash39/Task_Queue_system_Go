// Package for our Queue
package queue

import (
	"jobqueue/internal/task" //Importing the task package to use the Task struct defined in task.go
	"sync"                   // Importing sync package for mutex to ensure thread safety
)

type Queue struct {
	tasks []task.Task // Slice to hold our tasks
	mu    sync.Mutex  // mutex to ensure that the queue operations are thread-safe
}

func (q *Queue) Enqueue(t task.Task) { //we use task.Task struct to define the type of task we want to add to our queue
	q.mu.Lock() //Lock the mutex before modyfying the Queue to ensure that no other goroutine can access it at the same time

	defer q.mu.Unlock() //Unlock the mutex after the function returns, ensuring that it will be unlocked even if an error occurs

	q.tasks = append(q.tasks, t) //Adds task to the end of the tasks slice, effectively adding it to the queue
}

func (q *Queue) Dequeue() (task.Task, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.tasks) == 0 { //check if the queue is empty, if it is, we return an empty task and false to indicate that there was no task to dequeue
		return task.Task{}, false
	}

	t := q.tasks[0] //Gets the first task in the queue and stores it in the variable t

	q.tasks = q.tasks[1:] //removes the first tasks(deque) by slicing tasks to exclude the first element

	return t, true
}


