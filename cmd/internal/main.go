package main

import (
	"fmt"
	"smart-board/routes"
)

func main() {
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates"))))
	// http.HandleFunc("/", routes.MainHandler)
	// http.HandleFunc("/game", routes.GameHandler)
	// http.HandleFunc("/survey",routes.SurveyHandler)
	// fmt.Println("Server starts at: http://localhost:8080/")
	// http.ListenAndServe(":8080", nil)
	r := routes.InitializeRoutes()
	fmt.Println("Server starts at: http://localhost:8080/")
	r.Run(":8080")
	
	
}
