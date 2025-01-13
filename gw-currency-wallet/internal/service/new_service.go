package service

import (
	"wallet/api/grpc"
	"wallet/internal/storages"
)

func NewService() *Service {
	rep := storages.CreateRepository()
	rep.PrepareTable()
	return &Service{
		Repository: rep,
		GRPCClient: grpc.NewClient(),
	}
}
