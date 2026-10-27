package httpapi

import (
	"errors"
	"fmt"
	"net/http"
)

type Server struct {
	httpHandlers *Handler
}

func NewServer(httpHandlers *Handler) *Server {
	return &Server{
		httpHandlers: httpHandlers,
	}
}

func (hs *Server) Start() error {

	// Маршрутизатор от net/http, стоковый с go 1.22
	router := http.NewServeMux()

	// Заполняем маршрутизатор обработчиками
	router.HandleFunc("POST /game/start}", hs.httpHandlers.startGame)

	// Более корректное создание сервера, тут можно различных финтифлюшек добавить
	server := &http.Server{
		Addr:    ":9091",
		Handler: router,
	}

	fmt.Println("starting HTTP server on :9091")
	err := server.ListenAndServe()

	// Если сервер запустился с ошибкой или если завершился не штатно, http.ErrServerClosed - ошибка, в случае если сервер штатно завершил свою работу
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
