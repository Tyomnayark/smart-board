package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Здесь вы можете установить политику проверки Origin, если это необходимо.
		// Например, вы можете разрешить только запросы с определенных доменов.
		return true
	},
}

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

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Ошибка при установке WebSocket-соединения:", err)
		return
	}
	defer conn.Close()

	for {
		// Чтение сообщения от клиента
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Ошибка при чтении сообщения:", err)
			return
		}

		// Обработка сообщения, например, отправка его другим клиентам
		// В данном примере, просто отправляем обратно то же сообщение
		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println("Ошибка при отправке сообщения:", err)
			return
		}
	}
}
