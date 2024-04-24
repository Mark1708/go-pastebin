package router

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	hh "github.com/Mark1708/go-pastebin/internal/handler/healthcheck"
	ph "github.com/Mark1708/go-pastebin/internal/handler/paste"
	"github.com/Mark1708/go-pastebin/pkg/logger"
	"github.com/Mark1708/go-pastebin/pkg/rest"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

const (
	// CorsMaxAge Максимальное значение, не игнорируемое ни одним из основных браузеров.
	CorsMaxAge = 300
)

func New(
	pasteHandler ph.Handler,
	healthHandler hh.Handler,

) *chi.Mux {
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
		MaxAge:           CorsMaxAge,
	}))

	r.Route("/api/v1", func(r chi.Router) {
		rest.Get(r, "/health", healthHandler.CheckHealth)

		r.Route("/pastes", func(r chi.Router) {
			rest.Post(r, "/", pasteHandler.CreatePaste)
			r.Route("/{hash}", func(r chi.Router) {
				rest.Get(r, "/", pasteHandler.GetPaste)
				rest.Put(r, "/", pasteHandler.UpdatePaste)
				rest.Delete(r, "/", pasteHandler.DeletePaste)
			})
		})
	})

	workDir, _ := os.Getwd()

	// Конфигурация Swagger
	openAPIDir := http.Dir(filepath.Join(workDir, "api/openapi"))
	FileServer(r, "/api/v1/swagger", openAPIDir)
	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		logger.Log.Debug("FileServer does not permit any URL parameters.")
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
