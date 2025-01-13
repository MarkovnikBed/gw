package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	pb "github.com/MarkovnikBed/exchange_grpc/proto"
)

// @Summary      Обмен валюты
// @Description  Обмен
// @Tags         rates
// @Accept       json
// @Produce      json
// @Param        user_data body exchange true "Данные пользователя"
// @Param        Authorization header string true "токен авторизации"
// @Success      200 {object} response "обмен осуществлён"
// @Failure      400 {object} response "неверная валюта"
// @Router       /api/v1/exchange [post]
func (h *Handler) Exchange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()
	request := exchange{}
	json.NewDecoder(r.Body).Decode(&request)
	resp, err := h.Service.GRPCClient.GRPCClient.GetExchangeRateForCurrency(context.Background(), &pb.CurrencyRequest{
		FromCurrency: request.From,
		ToCurrency:   request.To,
	})
	if err != nil {
		json.NewEncoder(w).Encode(response{Error: err.Error()})
		return
	}

	token := r.Header.Get("Authorization")
	USD, EUR, RUB, err := h.Service.Repository.Exchange(token, resp.FromCurrency, resp.ToCurrency, float64(resp.Rate), request.Amount)
	if err != nil {
		json.NewEncoder(w).Encode(&response{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(&response{Message: "обмен осуществлён", NewBalance: &Rates{USD: float32(USD), EUR: float32(EUR), RUB: float32(RUB)}})
}
