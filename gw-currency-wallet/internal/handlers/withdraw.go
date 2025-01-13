package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// @Summary      Получение средств
// @Description  Получение СРЕДСТВ
// @Tags         operations
// @Accept       json
// @Produce      json
// @Param        user_data body AccountChange true "Данные пользователя" Example({"amount": 100, "currency": "USD"})
// @Param        Authorization header string true "токен авторизации"
// @Success      200 {object} response "cредства получены"
// @Failure      400 {object} response "неверная валюта"
// @Router       /api/v1/wallet/withdraw [post]
func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	query := AccountChange{}
	json.NewDecoder(r.Body).Decode(&query)
	log.Println(query.Amount)
	if query.Amount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&response{Error: "неверное количество средств"})
		return
	}
	switch query.Currency {
	case "USD", "RUB", "EUR":
	default:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&response{Error: "неверная валюта"})
		return
	}
	token := r.Header.Get("Authorization")
	USD, EUR, RUB, err := h.Service.Repository.Withdraw(token, query.Amount, query.Currency)
	if err != nil {
		json.NewEncoder(w).Encode(&response{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(&response{Message: "средства получены", NewBalance: &Rates{USD: float32(USD), EUR: float32(EUR), RUB: float32(RUB)}})
}
