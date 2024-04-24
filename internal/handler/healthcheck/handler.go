package healthcheck

import (
	"context"
	"net/http"

	"github.com/Mark1708/go-pastebin/pkg/operation"
	"github.com/Mark1708/go-pastebin/pkg/rest"
)

type Handler interface {
	CheckHealth(
		ctx context.Context,
		r *http.Request,
	) (rest.ResponseHolder[operation.ResponseDto], error)
}

type handler struct{}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) CheckHealth(_ context.Context, _ *http.Request) (
	rest.ResponseHolder[operation.ResponseDto],
	error,
) {
	return rest.StatusOk(operation.ResponseDto{Message: "Healthy"}), nil
}
