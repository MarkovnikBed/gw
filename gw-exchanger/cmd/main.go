package main

import (
	"log"
	"net"
	"os"

	pb "github.com/MarkovnikBed/exchange_grpc/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"server/internal/redis"
	"server/internal/server"
)

func init() {
	if err := godotenv.Load("config.env"); err != nil {
		log.Fatal("не удалось подгрузить конфигурационный файл")
	}
	redis.StartRedis()
}

func main() {
	port := os.Getenv("GRPC_PORT")
	list, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("не получилось прослушать порт № %v: %v", port, err)
	}
	s := grpc.NewServer()

	pb.RegisterExchangeServiceServer(s, &server.Server{})
	err = s.Serve(list)
	if err != nil {
		log.Fatalf("не удалось создать сервер: %v", err)
	}
}
