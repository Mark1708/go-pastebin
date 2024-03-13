package handler

import (
	"net/http"
	"strings"

	"github.com/Mark1708/go-pastebin/internal/paste"
	dto2 "github.com/Mark1708/go-pastebin/internal/paste/dto/request"

	"github.com/go-chi/render"

	hasher "github.com/Mark1708/go-pastebin/internal/pkg/hash"
)

type API struct {
	Service paste.Service
}

func (a *API) Get(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")

	paste, getErr := a.Service.GetByHash(hash)
	if !getErr.IsEmpty() {
		_ = render.Render(w, r, getErr)
		return
	}

	_ = render.Render(w, r, paste)
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	dto := &dto2.Dto{}
	_ = render.Bind(r, dto)
	ip := hasher.GetUserIP(r)

	createdDto, createError := a.Service.Create(ip, *dto)
	if !createError.IsEmpty() {
		_ = render.Render(w, r, createError)
		return
	}

	_ = render.Render(w, r, createdDto)
}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")

	dto := &dto2.Dto{}

	_ = render.Bind(r, dto)
	updatedDto, updateErr := a.Service.Update(hash, *dto)
	if !updateErr.IsEmpty() {
		_ = render.Render(w, r, updateErr)
		return
	}

	_ = render.Render(w, r, updatedDto)
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")

	// Игнорируем статичные файлы которые запрашивает браузер
	if strings.Contains(hash, ".") {
		http.NotFound(w, r)
		return
	}
	responseDto, deleteErr := a.Service.Delete(hash)
	if !deleteErr.IsEmpty() {
		_ = render.Render(w, r, deleteErr)
		return
	}

	_ = render.Render(w, r, responseDto)
}
