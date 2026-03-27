package main //main package that run thes go program, it is the entry point of the application

//importing the external packages and the one we made in the project
import (
	"fmt"
	
	"jobqueue/internal/queue"
	"jobqueue/internal/worker"
	"jobqueue/internal/api"
	"net/http"
	
)

func main() {
	q := &queue.Queue{}

	// Create API Server
	apiServer := &api.APIServer{Queue: q}

	// Register the route
	http.HandleFunc("/tasks", apiServer.HandleSubmitTask)

	// Start the server...

	for i:= 1; i<=3;i++{
		// Create worker
		w := worker.Worker{
			ID:    i,
			Queue: q,  //sharing the same queue on which they are working, this allows them to process tasks from the same queue concurrently
		}

		// 🔥 Start worker in background
		go w.Start()
	}
	
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}