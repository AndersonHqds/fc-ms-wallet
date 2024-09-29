package webserver

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webserverPort string) *WebServer {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return &WebServer{
		Router:        r,
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: webserverPort,
	}
}

func (ws *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	ws.Handlers[path] = handler
	ws.Router.HandleFunc(path, handler)
}

func (ws *WebServer) Start() error {
	for path, handler := range ws.Handlers {
		ws.Router.Post(path, handler)
	}
	return http.ListenAndServe(ws.WebServerPort, ws.Router)
}
