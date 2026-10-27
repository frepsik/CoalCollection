package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
)

// Вспомогательная функция для формирования ответов с ошибкой в качестве json (функция helper)
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json") //Дословно говорим, что в body запроса, будет именно json
	w.WriteHeader(status)

	// json.NewEncoder - говорим, что ответ будет положен напрямую в объект, который формирует ответ клиенту (w).
	// Encode(data) - функция, которая принимаемые данные, приведёт к типу json
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to write JSON response: %v", err)
	}
}
