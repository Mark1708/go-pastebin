package main

import (
	"io"
	"log"
	"net/http"
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

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Hello, world!")
	if err != nil {
		panic(err)
	}
	log.Printf("Host: %s", req.Host)
}
