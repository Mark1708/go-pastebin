package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Mark1708/go-pastebin/internal/config"
)

func main() {
	// Парсим конфигурацию
	c := config.New()

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
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      mux,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server " + s.Addr)
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
