package main

import (
	"os"
	"path/filepath"
	"smart-board/routes"

	"net/http"
)

func GetProjectRoot() string {
	projectRoot, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return projectRoot
}

func main() {
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates"))))
	// http.HandleFunc("/", routes.MainHandler)
	// http.HandleFunc("/game", routes.GameHandler)
	// http.HandleFunc("/survey",routes.SurveyHandler)
	// fmt.Println("Server starts at: http://localhost:8080/")
	// http.ListenAndServe(":8080", nil)

	r := routes.InitializeRoutes()
	websocketRouter := routes.CreateWebSocketRouter()

    go func() {
        if err := r.Run(":8080"); err != nil {
            panic(err)
        }
    }()

    http.Handle("/ws", websocketRouter) 

    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
