package queue

import(
	"sync"
	"jobqueue/internal/task"
)

type Queue struct {
	tasks []task.Task
	mu sync.Mutex
}

func (q *Queue) Enqueue(t task.Task) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.tasks = append(q.tasks, t)
}

func (q *Queue) Dequeue() (task.Task, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.tasks) == 0 {
		return task.Task{}, false
	}

	t := q.tasks[0]
	q.tasks = q.tasks[1:]

	return t, true
}