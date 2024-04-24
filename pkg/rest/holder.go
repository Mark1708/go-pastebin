package rest

import "net/http"

type ResponseHolder[T any] struct {
	Data   T
	Status int
}

func StatusOk[T any](data T) ResponseHolder[T] {
	return ResponseHolder[T]{Data: data, Status: http.StatusOK}
}

func StatusCreated[T any](data T) ResponseHolder[T] {
	return ResponseHolder[T]{Data: data, Status: http.StatusCreated}
}
