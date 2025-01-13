package server

import (
	"context"
	"strconv"

	pb "github.com/MarkovnikBed/exchange_grpc/proto"

	red "server/internal/redis"
	"server/internal/rpc"
)

func (s *Server) GetExchangeRates(ctx context.Context, r *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	redisClient := red.GetRedisClient()
	result, _ := redisClient.Client.HGetAll(context.Background(), "currency").Result()

	if len(result) == 0 {

		return rpc.GetExchangeRates(redisClient)
	}
	USD, _ := strconv.ParseFloat(result["USD"], 32)
	EUR, _ := strconv.ParseFloat(result["EUR"], 32)
	RUB, _ := strconv.ParseFloat(result["RUB"], 32)
	return &pb.ExchangeRatesResponse{
		Rates: map[string]float32{
			"USD": float32(USD),
			"EUR": float32(EUR),
			"RUB": float32(RUB),
		},
	}, nil

}
