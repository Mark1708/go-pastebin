package paste

import "net/http"

type API struct{}

func (a *API) Get(w http.ResponseWriter, r *http.Request) {}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {}
