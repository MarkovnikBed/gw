package grpc

import (
	"log"
	"os"

	pb "github.com/MarkovnikBed/exchange_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() Client {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		log.Fatal("в конфигурационном файле нет данных о переменной GRPC_PORT")
	}
	conn, err := grpc.NewClient(os.Getenv("GRPC_HOST")+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не удалось установить соединение по gRPC c портом %v", err)
	}
	client := pb.NewExchangeServiceClient(conn)

	return Client{
		GRPCClient: client,
	}
}
