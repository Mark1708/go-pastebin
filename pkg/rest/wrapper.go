package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Mark1708/go-pastebin/pkg/gerror"
	"github.com/go-chi/chi/v5"
)

func JSONWrapperHandler[T any](f func(http.ResponseWriter, *http.Request) (ResponseHolder[T], error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var response ResponseHolder[T]
		response, err = f(w, r)
		if err == nil {
			var jsonBytes []byte
			jsonBytes, err = json.MarshalIndent(response.Data, "", " ")
			if err == nil {
				w.WriteHeader(response.Status)
				_, _ = w.Write(jsonBytes)
				return
			}
		}
		if err != nil {
			gerror.Handle(w, err)
		}
	}
}

func JSONWrapperHandlerWithContext[T any](
	f func(ctx context.Context, r *http.Request) (ResponseHolder[T], error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var err error
		var response ResponseHolder[T]
		response, err = f(ctx, r)
		if err == nil {
			var jsonBytes []byte
			jsonBytes, err = json.MarshalIndent(response.Data, "", " ")
			if err == nil {
				w.WriteHeader(response.Status)
				_, _ = w.Write(jsonBytes)
				return
			}
		}
		if err != nil {
			gerror.Handle(w, err)
		}
	}
}

func Get[T any](
	r chi.Router,
	pattern string,
	f func(ctx context.Context, r *http.Request) (ResponseHolder[T], error),
) {
	r.Get(pattern, JSONWrapperHandlerWithContext(f))
}

func Post[T any](
	r chi.Router,
	pattern string,
	f func(ctx context.Context, r *http.Request) (ResponseHolder[T], error),
) {
	r.Post(pattern, JSONWrapperHandlerWithContext(f))
}

func Put[T any](
	r chi.Router,
	pattern string,
	f func(ctx context.Context, r *http.Request) (ResponseHolder[T], error),
) {
	r.Put(pattern, JSONWrapperHandlerWithContext(f))
}

func Delete[T any](
	r chi.Router,
	pattern string,
	f func(ctx context.Context, r *http.Request) (ResponseHolder[T], error),
) {
	r.Delete(pattern, JSONWrapperHandlerWithContext(f))
}
