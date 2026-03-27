package worker

import(
	"fmt"
	"time"
	"jobqueue/internal/queue"
	"jobqueue/internal/task"
)

//worker struct that will process the tasks from the queue, it has an ID to identify the worker and a reference to the queue it will be processing tasks from
type Worker struct{
	ID int
	Queue *queue.Queue
}

func (w *Worker) Start() {
	for {
		//Taking a task from the queue using the Dequeue method, which returns a task and a boolean indicating whether a task was successfully dequeued. If no task is available, it waits for a second before trying again.
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