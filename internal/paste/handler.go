package paste

import (
	"net/http"
)

type API struct{}

func (a *API) Get(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Ok"))
}

func (a *API) Create(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Ok"))
}

func (a *API) Update(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Ok"))
}

func (a *API) Delete(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Ok"))
}
