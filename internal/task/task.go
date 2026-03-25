// This creates the package for our project file
package task

// Define a type for task status
type Status string

// Define possible statuses for tasks

const (
	Pending    Status = "pending"
	Processing Status = "processing"
	Completed  Status = "completed"
	Failed     Status = "failed"
)

// There goes our structure for every task
type Task struct {
	ID     string
	Type   string
	Data   string
	Status Status
}
