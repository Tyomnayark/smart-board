package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"html/template"
)

type News struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
	Date  int    `json:"date"`
	Text  string `json:"text"`
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	var news []News

	file, err := os.Open("smart-board/news.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	newsJson, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	err = json.Unmarshal(newsJson, &news)
	if err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		return
	}

	tmpl, err := template.ParseFiles("./assets/index.html")
	if err != nil {
		fmt.Println("3а:", err)
		return
	}

	err = tmpl.Execute(w, news)
	if err != nil {
		fmt.Println("4:", err)
		return
	}
}
