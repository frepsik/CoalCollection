package httpapi

// Проверяем, корректное ли значение пришло для запуска/завершения работы предприятия
// func validate() error {
// 	state := strings.TrimSpace(g.state)
// 	if state == "" {
// 		return errors.New("Argument operation game is empty")
// 	}
// 	if !(state == "start" || state == "shutdown") {
// 		return errors.New("Invalid operation game")
// 	}
// 	return nil
// }

type errorResponse struct {
	Message string `json:"message"`
}

func writeError(message string) errorResponse {
	return errorResponse{
		Message: message,
	}
}
