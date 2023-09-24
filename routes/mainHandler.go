package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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

func GetProjectRoot() string {
	projectRoot, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		panic(err)
	}
	return projectRoot

	// dir, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// return dir

}

func MainHandler(c *gin.Context) {
	var news []News

	// currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	filePath := filepath.Join(GetProjectRoot() + "/news.json")

	file, err := os.Open(filePath)
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

	tmpl, err := template.ParseFiles(GetProjectRoot() + "/assets/index.html")
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
	// dir := GetProjectRoot()

	r := gin.Default()
	r.Static("/assets/", "./assets")
	r.GET("/", MainHandler)
	r.GET("/survey", SurveyHandler)
	r.GET("/api/:id", getItem)
	r.GET("/api/", GetAll)
	r.POST("/api", createItem)
	r.PUT("/api/:id", updateItem)
	r.DELETE("api/:id", deleteItem)
	return r
}
