package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	ReadTimeout  = 2 // секунд
	WriteTimeout = 2 // секунд
	// IdleTimeout - это максимальное время ожидания следующего запроса при включении функции keep-alives.
	IdleTimeout = 5 // секунд
)

func main() {
	/*
	 * ServeMux - это мультиплексор HTTP-запросов. Он сопоставляет URL каждого входящего запроса
	 * со списком зарегистрированных шаблонов и вызывает обработчик шаблона,
	 * который наиболее точно соответствует URL. NewServeMux выделяет и возвращает новый ServeMux.
	 *
	 * Поскольку ServeMux по умолчанию очень ограничен и не очень производителен
	 */
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  ReadTimeout * time.Second,
		WriteTimeout: WriteTimeout * time.Second,
		IdleTimeout:  IdleTimeout * time.Second,
	}

	log.Println("Starting server :8080")
	// Сервер определяет параметры для запуска HTTP-сервера.
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server startup failed")
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Hello, world!")
	if err != nil {
		panic(err)
	}
	log.Printf("Host: %s", req.Host)
}
