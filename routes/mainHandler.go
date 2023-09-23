package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type News struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
	Date  int    `json:"date"`
	Text  string `json:"text"`
}

func MainHandler(c *gin.Context) {
	var news []News

	file, err := os.Open("news.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при открытии файла"})
		return
	}
	defer file.Close()

	newsJson, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при чтении файла"})
		return
	}

	err = json.Unmarshal(newsJson, &news)
	if err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при парсинге JSON"})
		return
	}

	tmpl, err := template.ParseFiles("./assets/index.html")
	if err != nil {
		fmt.Println("3а:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при загрузке HTML шаблона"})
		return
	}

	data := gin.H{
		"news": news,
	}

	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		fmt.Println("4:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении шаблона"})
		return
	}
}

func InitializeRoutes() *gin.Engine {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Static("/assets", dir+"/assets")
	r.GET("/", MainHandler)
	r.GET("/survey", SurveyHandler)
	r.GET("/game",GameHandler)

	return r
}
