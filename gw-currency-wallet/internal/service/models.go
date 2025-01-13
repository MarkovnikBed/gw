package service

import (
	"wallet/api/grpc"
	"wallet/internal/storages"
)

type Service struct {
	Repository storages.Storage
	GRPCClient grpc.Client
}
