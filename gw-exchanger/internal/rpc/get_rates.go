package rpc

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	pb "github.com/MarkovnikBed/exchange_grpc/proto"

	"server/internal/redis"
)

func GetExchangeRates(rc *redis.RedisClient) (*pb.ExchangeRatesResponse, error) {
	rates := make(map[string]float32)
	resp, err := http.Get("https://www.cbr-xml-daily.ru/latest.js")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := &response{}

	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		return nil, err
	}

	rates["EUR"] = 1 / r.Rates.EUR
	rates["USD"] = 1 / r.Rates.USD
	rates["RUB"] = 1
	if err := rc.Client.HSet(context.Background(), "currency", "USD", rates["USD"], "EUR", rates["EUR"], "RUB", rates["RUB"]).Err(); err != nil {
		log.Println(err)
	}
	rc.Client.Expire(context.Background(), "currency", 20*time.Second)
	return &pb.ExchangeRatesResponse{
		Rates: rates,
	}, nil

}
