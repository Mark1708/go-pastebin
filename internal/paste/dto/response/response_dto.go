package response

import (
	"net/http"

	"github.com/Mark1708/go-pastebin/internal/paste/dto/visibility"
)

type Dto struct {
	Hash       string         `json:"hash"`
	Title      string         `json:"title"`
	Visibility visibility.Dto `json:"visibility"`
	Content    string         `json:"content"`
	CreatedAt  string         `json:"created_at"`
	ExpiredAt  string         `json:"expired_at"`
}

func (rsDTO Dto) Render(w http.ResponseWriter, _ *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}
