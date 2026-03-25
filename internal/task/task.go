// This creates the package for our project file
package task

type Status string


const (
	Pending Status ="pending"
	Processing Status = "processing"
	completed Status = "completed"
	Failed Status = "failed"
)

// There goes our structure for every task
type Task struct {
	ID string
	Type string
	Data string
	Status Status
}

