package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Mark1708/go-pastebin/internal/router"

	"github.com/Mark1708/go-pastebin/internal/config"
)

//  @title          Pastebin API
//  @version        1.0
//  @description    This is Pastebin example with Golang

//  @contact.name   Mark
//  @contact.url    https://github.com/Mark1708
// 	@contact.email  mark1708.work@gmail.com

//  @license.name   MIT License
//  @license.url    LICENSE

// @host       localhost:8080
// @basePath   /api
// @schemes	   http https
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
