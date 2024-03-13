package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Mark1708/go-pastebin/internal/router"

	"github.com/Mark1708/go-pastebin/internal/config"
)

func main() {
	// Парсим конфигурацию
	c := config.New()

	// Инициализируем роутер
	r := router.New()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
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
