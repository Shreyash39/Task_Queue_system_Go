package api
import(
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"jobqueue/internal/queue"
	"jobqueue/internal/task"

)

type APIServer struct{
	Queue *queue.Queue
}

func (s* APIServer)HandleSubmitTask(w http.ResponseWriter, r *http.Request){	
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var t task.Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	t.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	t.Status = task.Pending

	s.Queue.Enqueue(t)

	fmt.Fprintf(w, "Task %s submitted successfully", t.ID)
}