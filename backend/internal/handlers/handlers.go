package handlers

import (
	"html/template"
	"net/http"
)

// IndexHandler обрабатывает запросы к корневому маршруту
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Загружаем шаблон
	tmpl, err := template.ParseFiles("../../Public/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Данные для передачи в шаблон
	data := struct {
		Title string
		CSS1  string
	}{
		Title: "Добро пожаловать на главную страницу!",
		CSS1:  "/assets/styles/style.css",
	}

	// Выполняем шаблон и отправляем результат в ResponseWriter
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
