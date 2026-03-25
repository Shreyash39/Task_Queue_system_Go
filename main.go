package main

import (
	"fmt"
	"jobqueue/internal/queue"
	"jobqueue/internal/task"
)

func main() {
	q := &queue.Queue{}

	// Create sample task
	t := task.Task{
		ID:     "1",
		Type:   "email",
		Data:   "Send welcome email",
		Status: task.Pending,
	}

	// Add to queue
	q.Enqueue(t)

	fmt.Println("Task added to queue")

	// Remove from queue
	taskOut, ok := q.Dequeue()
	if ok {
		fmt.Println("Dequeued Task:", taskOut)
	} else {
		fmt.Println("Queue is empty")
	}
}