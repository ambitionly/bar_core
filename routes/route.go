package routes

import (
	"log"
	"net/http"
	"os"
)

// Dispatch 路由分发
func Dispatch() {
	http.HandleFunc("/api/post", controllers.AddPost)

	err := http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("LISTEN"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
