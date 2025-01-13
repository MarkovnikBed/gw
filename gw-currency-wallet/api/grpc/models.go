package grpc

import (
	pb "github.com/MarkovnikBed/exchange_grpc/proto"
)

type Client struct {
	GRPCClient pb.ExchangeServiceClient
}
