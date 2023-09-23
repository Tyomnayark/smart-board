package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GameHandler(c *gin.Context) {
	tmpl, err := template.ParseFiles("./templates/assets/gamepage.html")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка при загрузке шаблона: %s", err))
		return
	}

	if err := tmpl.Execute(c.Writer, nil); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка при выполнении шаблона: %s", err))
		return
	}
}
