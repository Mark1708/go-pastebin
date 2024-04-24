package paste

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Mark1708/go-pastebin/pkg/ip"

	model "github.com/Mark1708/go-pastebin/internal/model/paste"
	service "github.com/Mark1708/go-pastebin/internal/service/paste"
	"github.com/Mark1708/go-pastebin/pkg/logger"
	"github.com/Mark1708/go-pastebin/pkg/operation"
	"github.com/Mark1708/go-pastebin/pkg/rest"
	"go.uber.org/zap"
)

type Handler interface {
	GetPaste(
		ctx context.Context,
		r *http.Request,
	) (rest.ResponseHolder[model.ResponseDto], error)
	CreatePaste(
		ctx context.Context,
		r *http.Request,
	) (rest.ResponseHolder[model.ResponseDto], error)
	UpdatePaste(
		ctx context.Context,
		r *http.Request,
	) (rest.ResponseHolder[model.ResponseDto], error)
	DeletePaste(
		ctx context.Context,
		r *http.Request,
	) (rest.ResponseHolder[operation.ResponseDto], error)
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetPaste(
	ctx context.Context,
	r *http.Request,
) (rest.ResponseHolder[model.ResponseDto], error) {
	hash := r.PathValue("hash")

	dto, err := h.service.GetPasteByHash(ctx, hash)
	return rest.StatusOk(dto), err
}

func (h *handler) CreatePaste(
	ctx context.Context,
	r *http.Request,
) (rest.ResponseHolder[model.ResponseDto], error) {
	ctx = context.WithValue(ctx, ip.UserIPKey{}, ip.GetUserIP(r))

	reqDto := &model.RequestDto{}

	err := json.NewDecoder(r.Body).Decode(&reqDto)
	if err != nil {
		logger.Log.With(zap.Error(err)).Error("error during read request dto")
		return rest.ResponseHolder[model.ResponseDto]{}, err
	}

	dto, err := h.service.CreatePaste(ctx, *reqDto)
	return rest.StatusCreated(dto), err
}

func (h *handler) UpdatePaste(
	ctx context.Context,
	r *http.Request,
) (rest.ResponseHolder[model.ResponseDto], error) {
	hash := r.PathValue("hash")

	reqDto := &model.RequestDto{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&reqDto)
	if err != nil {
		logger.Log.With(zap.Error(err)).Error("error during read request dto")
		return rest.ResponseHolder[model.ResponseDto]{}, err
	}

	dto, err := h.service.UpdatePaste(ctx, hash, *reqDto)
	return rest.StatusOk(dto), err
}

func (h *handler) DeletePaste(
	ctx context.Context,
	r *http.Request,
) (rest.ResponseHolder[operation.ResponseDto], error) {
	hash := r.PathValue("hash")

	return rest.StatusOk(operation.ResponseDto{Message: "Deleted"}), h.service.DeletePaste(ctx, hash)
}
