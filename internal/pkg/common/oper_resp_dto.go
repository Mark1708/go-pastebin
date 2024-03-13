package common

import "net/http"

type OperationResponseDto struct {
	Message string `json:"message"`
}

func (orDTO OperationResponseDto) Render(w http.ResponseWriter, _ *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}
