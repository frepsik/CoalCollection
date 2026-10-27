package http

import (
	"errors"
	"fmt"
	"net/http"
)

type HttpServer struct {
	httpHandlers *HttpHandlers
}

func NewHttpServer(httpHandlers *HttpHandlers) *HttpServer {
	return &HttpServer{
		httpHandlers: httpHandlers,
	}
}

func (hs *HttpServer) StartServer() error {

	// Маршрутизатор от net/http, стоковый с go 1.22
	router := http.NewServeMux()

	// Заполняем маршрутизатор обработчиками
	router.HandleFunc("POST /game/start}", hs.httpHandlers.handleStartGame)

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
