package routes

import (
	"net/http"

	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Можете настроить проверку origin, если необходимо
		return true
	},
}

func CreateWebSocketRouter() http.Handler {
	http.Handle("/ws", http.HandlerFunc(handleWebSocketConnection))

	return http.DefaultServeMux
}

func handleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	// Попытка обновления HTTP-запроса до WebSocket-соединения
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// Обработка ошибки при обновлении
		return
	}
	defer conn.Close()

	// Обработка WebSocket-соединения
	for {
		// Чтение сообщения от клиента
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			// Обработка ошибки чтения
			return
		}

		// Обработка сообщения, например, отправка его обратно клиенту
		if err := conn.WriteMessage(messageType, p); err != nil {
			// Обработка ошибки отправки
			return
		}
	}
}
