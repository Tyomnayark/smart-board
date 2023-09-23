package main

import (
	"fmt"
	"net/http"
	"smart-board/routes"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates"))))
	http.HandleFunc("/", routes.MainHandler)
	// http.HandleFunc("/game", routes.gameHandler)
	// http.HandleFunc("/map", routes.mapHandler)
	fmt.Println("Server starts at: http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
