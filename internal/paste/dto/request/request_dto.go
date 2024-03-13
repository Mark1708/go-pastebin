package request

import (
	"net/http"
)

type Dto struct {
	Title      string `json:"title"`
	Visibility string `json:"visibility"`
	Content    string `json:"content"`
	Expires    string `json:"expires"`
}

func (rqDTO Dto) Bind(_ *http.Request) error {
	return nil
}
