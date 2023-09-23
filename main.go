package main

import (
	"log"
	"os"
	"path/filepath"
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
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	r := routes.InitializeRoutes()

	// t := gin.Default()

	r.Static("/assets", dir+"/assets")

	// t.LoadHTMLGIob(dir + "/templates/*")

	// r.GET ("/", controllers. NotesIndex)
	// r.POST ("/", controllers. NoteCreate)

	// Start app

	r.Run(":8080")

}
