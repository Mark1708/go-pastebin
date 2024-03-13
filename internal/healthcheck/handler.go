package healthcheck

import (
	"net/http"

	"github.com/Mark1708/go-pastebin/internal/pkg/common"
	"github.com/go-chi/render"
)

func Check(w http.ResponseWriter, r *http.Request) {
	health := common.OperationResponseDto{Message: "Healthy"}
	_ = render.Render(w, r, health)
}
