package gerror

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mark1708/go-pastebin/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Handle(_ http.ResponseWriter, err error) {
	// TODO: Закончить с обработчиком ошибок
	switch err.(type) { //nolint:errorlint // Так надо
	case *json.UnmarshalTypeError:
		var typeError *json.UnmarshalTypeError
		errors.As(err, &typeError)
		logger.Log.With(zap.Error(err)).Errorf(
			"decode error of field %s (expected - \"%s\", but get \"%s\")",
			typeError.Field, typeError.Type.String(), typeError.Value,
		)
	case *json.SyntaxError:
		var typeError *json.SyntaxError
		errors.As(err, &typeError)
		logger.Log.With(zap.Error(err)).Error(
			"decode error of dto",
		)
		// writeError(w, BadArguments{Base{Status: 404, Code: "", Message: ""}})
	case *time.ParseError:
		var typeError *time.ParseError
		errors.As(err, &typeError)
		logger.Log.With(zap.Error(err)).Errorf(
			"incorrect datetime format (expected -  \"%s\", but get \"%s\")",
			typeError.Layout, typeError.Value,
		)
	default:
		switch err.Error() {
		case "unexpected EOF":
			logger.Log.With(zap.Error(err)).Error("incorrect json format")
		default:
			logger.Log.With(zap.Error(err)).Error("decode error")
		}
	}
}
