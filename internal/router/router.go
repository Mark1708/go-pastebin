package router

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Mark1708/go-pastebin/internal/paste/handler"
	"github.com/Mark1708/go-pastebin/internal/paste/repository"
	"github.com/Mark1708/go-pastebin/internal/paste/service"

	"github.com/jmoiron/sqlx"

	"github.com/Mark1708/go-pastebin/internal/config"

	"github.com/Mark1708/go-pastebin/internal/healthcheck"
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

	// Создаём подключение
	db, err := sqlx.Connect(
		"pgx",
		"postgresql://pb_user:pb_pass@db:5432/pastebin_db?sslmode=disable",
	)
	if err != nil {
		log.Fatalln(err)
	}

	r.Get(config.BasePath+"/healthcheck", healthcheck.Check)

	r.Route(config.BasePath, func(r chi.Router) {
		r.Route("/pastes", func(r chi.Router) {
			pasteRepo := &repository.Repository{DB: db}
			pasteService := &service.Service{Repo: *pasteRepo}
			pasteAPI := &handler.API{Service: pasteService}
			r.Post("/", pasteAPI.Create)
			r.Route("/{hash}", func(r chi.Router) {
				r.Get("/", pasteAPI.Get)
				r.Put("/", pasteAPI.Update)
				r.Delete("/", pasteAPI.Delete)
			})
		})
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
