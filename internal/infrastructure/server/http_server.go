package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"net/http"
)

// NewHTTPServer создает новый HTTP-сервер для GraphQL
func NewHTTPServer(schema graphql.ExecutableSchema) *http.Server {
	// Создание обработчика GraphQL
	srv := handler.NewDefaultServer(schema)

	// Настройка транспортов
	srv.AddTransport(transport.Websocket{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})

	// Создание маршрутизатора HTTP
	mux := http.NewServeMux()

	// Точка входа для GraphQL-запросов
	mux.Handle("/graphql", srv) // Обработчик GraphQL

	// Добавление GraphQL Playground для тестирования запросов
	mux.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

	// Настройка HTTP-сервера
	return &http.Server{
		Addr:    ":8080", // Порт, на котором будет запущен сервер
		Handler: mux,     // Обработчик HTTP-запросов
	}
}
