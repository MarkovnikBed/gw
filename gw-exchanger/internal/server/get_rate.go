package server

import (
	"context"
	"log"
	"strconv"

	pb "github.com/MarkovnikBed/exchange_grpc/proto"
	"github.com/go-redis/redis/v8"

	red "server/internal/redis"
	"server/internal/rpc"
)

func (s *Server) GetExchangeRateForCurrency(ctx context.Context, r *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
	redisClient := red.GetRedisClient()
	res, err1 := redisClient.Client.HGet(ctx, "currency", r.FromCurrency).Result()

	res2, err2 := redisClient.Client.HGet(ctx, "currency", r.ToCurrency).Result()

	if err1 == redis.Nil || err2 == redis.Nil {
		resp, _ := rpc.GetExchangeRates(redisClient)
		rate := resp.Rates[r.FromCurrency] / resp.Rates[r.ToCurrency]
		return &pb.ExchangeRateResponse{
			FromCurrency: r.FromCurrency,
			ToCurrency:   r.ToCurrency,
			Rate:         rate,
		}, nil
	}
	log.Println(res, res2)
	resFloat, _ := strconv.ParseFloat(res, 64)
	res2Float, _ := strconv.ParseFloat(res2, 64)
	rate := resFloat / res2Float
	return &pb.ExchangeRateResponse{
		FromCurrency: r.FromCurrency,
		ToCurrency:   r.ToCurrency,
		Rate:         float32(rate),
	}, nil
}
