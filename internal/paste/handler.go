package paste

import "net/http"

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)

	Create(w http.ResponseWriter, r *http.Request)

	Update(w http.ResponseWriter, r *http.Request)

	Delete(w http.ResponseWriter, r *http.Request)
}
