package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Survey struct {
	ID        int    `json:"id"`
	Image     string `json:"image"`
	Name      string `json:"name"`
	Date      int    `json:"date"`
	Questions string `json:"text"`
}

func SurveyHandler(c *gin.Context) {
	var survey Survey
	file, err := os.Open("question.json")
	defer file.Close()
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}

	surveyJson, err := io.ReadAll(file)

	err = json.Unmarshal(surveyJson, &survey)
	if err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		return
	}

	tmpl, err := template.ParseFiles("./templates/assets/index.html")
	if err != nil {
		fmt.Println("3а:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при загрузке HTML шаблона"})
		return
	}

	data := gin.H{
		"survey": survey,
	}

	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		fmt.Println("4:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении шаблона"})
		return
	}
}

func SurveyRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/survey", MainHandler)

	return r
}

func PostSurvey() *gin.Engine {
	r := gin.Default()

	r.POST("/survey", SurveyHandler )

	return r
}



// func SurveyHandler(w http.ResponseWriter, r *http.Request) {
// 	var survey Survey
// 	file, err := os.Open("question.json")
// 	defer file.Close()
// 	if err != nil {
// 		fmt.Println("Ошибка при открытии файла:", err)
// 		return
// 	}

// 	surveyJson, err := io.ReadAll(file)

// 	err = json.Unmarshal(surveyJson, &survey)
// 	if err != nil {
// 		fmt.Println("Ошибка при парсинге JSON:", err)
// 		return
// 	}

// 	tmpl, err := template.ParseFiles("./templates/assets/index.html")
// 	if err != nil {
// 		fmt.Println("3а:", err)
// 		return
// 	}

// 	err = tmpl.Execute(w, news)
// 	if err != nil {
// 		fmt.Println("4:", err)
// 		return
// 	}

// }
