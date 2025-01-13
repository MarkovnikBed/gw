package server

import (
	pb "github.com/MarkovnikBed/exchange_grpc/proto"
)

type Server struct {
	pb.ExchangeServiceServer
}
