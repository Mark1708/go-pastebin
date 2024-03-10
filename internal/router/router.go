package router

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Mark1708/go-pastebin/internal/config"

	"github.com/Mark1708/go-pastebin/internal/health"
	"github.com/Mark1708/go-pastebin/internal/paste"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	// Настройка CORS
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"},
		// AllowedOrigins: []string{"https://*", "http://*"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           config.CorsMaxAge,
	}))

	r.Get(config.BasePath+"/health", health.Check)

	r.Route(config.BasePath, func(r chi.Router) {
		pasteAPI := &paste.API{}
		r.Get("/pastes/{id}", pasteAPI.Get)
		r.Post("/pastes", pasteAPI.Create)
		r.Put("/pastes/{id}", pasteAPI.Update)
		r.Delete("/pastes/{id}", pasteAPI.Delete)
	})

	workDir, _ := os.Getwd()

	// Конфигурация Swagger
	openAPIDir := http.Dir(filepath.Join(workDir, "api/openapi-spec/build"))
	FileServer(r, config.BasePath+"/swagger", openAPIDir)

	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
