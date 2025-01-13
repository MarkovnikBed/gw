package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/MarkovnikBed/exchange_grpc/proto"
)

// @Summary      Получение курса валют
// @Description  Обмен
// @Tags         rates
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "токен авторизации"
// @Success      200 {object} Rates
// @Failure      500 {object} response "курсы валют на данный ммоент недоступны"
// @Router      /api/v1/exchange/rates [get]
func (h *Handler) GetRates(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	resp, err := h.Service.GRPCClient.GRPCClient.GetExchangeRates(ctx, &pb.Empty{})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Error: "курсы валют на данынй момент недоступны"})
		return
	}
	json.NewEncoder(w).Encode(&resp.Rates)

}
