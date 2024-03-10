package health

import (
	"net/http"
)

func Check(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Ok"))
}
