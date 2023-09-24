package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type Option string

// const {
// 	First Option = "multiple_choice"
// 	Second Option = "one_correct_an"
// }

type Question struct {
	QuestionText string   `json:"QuestionText"`
	QuestionType string   `json:"QuestionType"`
	Options      []string `json:"Options"`
}

type Questionary struct {
	ID          int        `json:"id"`
	SurveyTitle string     `json:"SurvetTitle"`
	Questions   []Question `json:"Questions"`
}

var questionary []Questionary

func ApiHandler(c *gin.Context) {
	var survey []Survey

	test := Questionary{
		ID:          1,
		SurveyTitle: "Customer Satisfaction Survey",
		Questions: []Question{
			{
				QuestionText: "How satisfied are you with our product?",
				QuestionType: "FirstOption",
				Options: []string{
					"Very Satisfied",
					"Satisfied",
					"Neutral",
					"Dissatisfied",
					"Very Dissatisfied",
				},
			},
		},
	}
	questionary = append(questionary, test)
	questionary = append(questionary, test)

	file, err := os.Open(GetProjectRoot() + "/question.json")
	defer file.Close()
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при открытии файла"})
		return
	}

	surveyJson, err := io.ReadAll(file)

	err = json.Unmarshal(surveyJson, &survey)
	if err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при парсинге JSON"})
		return
	}

	jsonData, err := json.Marshal(test)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(string(jsonData))

	c.JSON(http.StatusOK, test)

}

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, questionary)
}

func PostHandler(c *gin.Context) {

	var requestData map[string]interface{}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(requestData)

	c.JSON(http.StatusOK, gin.H{"message": "Данные успешно получены"})
}

func createItem(c *gin.Context) {
    var newItem Questionary
    if err := c.BindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newItem.ID = len(questionary) + 1
    questionary = append(questionary, newItem)

    c.JSON(http.StatusCreated, newItem)
}

func getItem(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	for _, item := range questionary {
		if int64(item.ID) == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func updateItem(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	var updatedItem Questionary
	if err := c.BindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, item := range questionary {
		if int64(item.ID) == id {
			updatedItem.ID = item.ID
			questionary[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func deleteItem(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	for i, item := range questionary {
		if int64(item.ID) == id {
			questionary = append(questionary[:i], questionary[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
