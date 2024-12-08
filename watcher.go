package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func watchFiles() {
	lastModTime := time.Now()

	for {
		time.Sleep(1 * time.Second)

		// Проверяем время изменения файла index.html
		info, err := os.Stat("ts/Public/index.html")
		if err != nil {
			log.Println("Ошибка при получении информации о файле:", err)
			continue
		}

		if info.ModTime().After(lastModTime) {
			lastModTime = info.ModTime()
			notifyClients()
		}
	}
}

var clients = make(map[*websocket.Conn]bool)

func notifyClients() {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte("reload"))
		if err != nil {
			log.Println("Ошибка при отправке сообщения клиенту:", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка при подключении WebSocket:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка при чтении сообщения:", err)
			delete(clients, conn)
			break
		}
	}
}
