package main //main package that run thes go program, it is the entry point of the application

//importing the external packages and the one we made in the project
import (
	"strconv"
	"time"
	
	"jobqueue/internal/queue"
	"jobqueue/internal/task"
	"jobqueue/internal/worker"
)

func main() {
	q := &queue.Queue{}

	// Create worker
	w := worker.Worker{
		ID:    1,
		Queue: q,
	}

	// 🔥 Start worker in background
	go w.Start()

	// Add tasks
	for i := 1; i <= 5; i++ {
		t := task.Task{
			ID:     strconv.Itoa(i),
			Type:   "email",
			Data:   "Send email",
			Status: task.Pending,
		}

		// Add to queue
		q.Enqueue(t)
	}

	// Keep program alive
	time.Sleep(10 * time.Second)
}