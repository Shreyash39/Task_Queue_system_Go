package worker

import(
	"fmt"
	"time"
	"jobqueue/internal/queue"
	"jobqueue/internal/task"
)

type Worker struct{
	ID int
	Queue *queue.Queue
}

func (w *Worker) Start() {
	for {
		t, ok := w.Queue.Dequeue()
		if !ok {
			time.Sleep(1 * time.Second)
			continue
		}

		fmt.Printf("Worker %d processing task %s\n", w.ID, t.ID)

		// simulate work
		time.Sleep(2 * time.Second)

		t.Status = task.Completed

		fmt.Printf("Worker %d completed task %s\n", w.ID, t.ID)
	}
}