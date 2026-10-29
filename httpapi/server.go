package httpapi

import (
	"errors"
	"fmt"
	"net/http"
)

type Server struct {
	handler *Handler
	server  *http.Server
	router  *http.ServeMux
}

func NewServer(handler *Handler) *Server {
	router := createRoute(handler)

	// Более корректное создание сервера, тут можно различных финтифлюшек добавить
	server := &http.Server{
		Addr:    ":9091",
		Handler: router,
	}

	return &Server{
		handler: handler,
		server:  server,
		router:  router,
	}
}

func createRoute(handler *Handler) *http.ServeMux {
	// Маршрутизатор от net/http, стоковый с go 1.22
	router := http.NewServeMux()

	// Заполняем маршрутизатор обработчиками
	router.HandleFunc("POST /game/start", handler.startGame)
	router.HandleFunc("POST /game/shutdown", handler.shutdownGame)
	router.HandleFunc("GET /game/status", handler.statusEnterprise)
	router.HandleFunc("GET /equipments/type", handler.equipmentsType)
	router.HandleFunc("GET /equipments", handler.equipments)
	router.HandleFunc("POST /equipments/{type}/purchase", handler.buyEquipment)
	router.HandleFunc("POST /miners", handler.hireMiner)
	router.HandleFunc("GET /miners/type", handler.minersType)
	router.HandleFunc("GET /miners", handler.miners)
	return router
}

func (s *Server) Start() error {

	fmt.Println("starting HTTP server on :9091")
	err := s.server.ListenAndServe()

	// Если сервер запустился с ошибкой или если завершился не штатно, http.ErrServerClosed - ошибка, в случае если сервер штатно завершил свою работу
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
