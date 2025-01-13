package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "wallet/docs"
	"wallet/internal/handlers"
	"wallet/internal/service"
)

func init() {

	if err := godotenv.Load("config.env"); err != nil {
		log.Fatalf("не удалось найти главный конфигурационный файл: %v", err)
	}
}

func main() {
	router := chi.NewRouter()

	service := service.NewService()
	handler := handlers.NewHandler(service)

	// @title Swagger WALLET API
	// @version 1.0
	// @description это простой кошелёк - обменник
	router.Post("/api/v1/register", handler.Register)
	router.Post("/api/v1/login", handler.Login)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	router.Post("/api/v1/wallet/deposit", handler.Auth(handler.Deposit))
	router.Post("/api/v1/wallet/withdraw", handler.Auth(handler.Withdraw))

	router.Get("/api/v1/exchange/rates", handler.Auth(handler.GetRates))
	router.Post("/api/v1/exchange", handler.Auth(handler.Exchange))

	start(router)
}

func start(router *chi.Mux) {
	port, ok := os.LookupEnv("HOST_PORT")
	if !ok || port == "" {
		log.Fatal("в конфигурационном файле отсутсвует переменная 'HOST_PORT'")
	}
	log.Printf("подготовка сервера на порту :%s...", port)
	list, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("не получилось прослушать порт: ", err)
	}

	server := http.Server{
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Println("сервер запущен ...")
	log.Fatal("прерывание работы сервера: ", server.Serve(list))
}
