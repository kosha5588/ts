package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	tasks []string
	mu    sync.Mutex
)

var upgrader = websocket.Upgrader{}
var clients = make(map[*websocket.Conn]bool)

func main() {
	go watchFiles() // Запускаем функцию отслеживания файлов

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/ws", handleWebSocket) // Обработка WebSocket

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("ts/Public/styles"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("ts/Public/scripts"))))

	log.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ts/Public/index.html")
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var newTask struct {
			Task string `json:"task"`
		}
		if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mu.Lock()
		tasks = append(tasks, newTask.Task)
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newTask)

	case http.MethodDelete:
		var taskToDelete struct {
			Task string `json:"task"`
		}
		if err := json.NewDecoder(r.Body).Decode(&taskToDelete); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mu.Lock()
		for i, task := range tasks {
			if task == taskToDelete.Task {
				tasks = append(tasks[:i], tasks[i+1:]...) // Удаляем задачу
				break
			}
		}
		mu.Unlock()

		w.WriteHeader(http.StatusNoContent) // Успешное удаление, без содержимого
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

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
