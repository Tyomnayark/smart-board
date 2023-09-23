package main

import (
	"smart-board/routes"
)

func main() {
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates"))))
	// http.HandleFunc("/", routes.MainHandler)
	// http.HandleFunc("/game", routes.GameHandler)
	// http.HandleFunc("/survey",routes.SurveyHandler)
	// fmt.Println("Server starts at: http://localhost:8080/")
	// http.ListenAndServe(":8080", nil)

	// Create gin app

	r := routes.InitializeRoutes()
	r.Run(":8080")
	// t := gin.Default()

	// r.Static("/assets", dir+"/assets")

	// t.LoadHTMLGIob(dir + "/templates/*")

	// r.GET ("/", controllers. NotesIndex)
	// r.POST ("/", controllers. NoteCreate)

	// Start app

}
