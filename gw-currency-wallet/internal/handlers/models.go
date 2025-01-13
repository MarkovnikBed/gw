package handlers

import "wallet/internal/service"

type Handler struct {
	Service *service.Service
}

type Rates struct {
	USD float32 `json:"USD"`
	EUR float32 `json:"EUR"`
	RUB float32 `json:"RUB"`
}

type UserData struct {
	Email    string `json:"email" example:"john@example.com"`
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"securepassword"`
}

type response struct {
	Message    string `json:"message,omitempty"`
	Error      string `json:"error,omitempty" `
	Token      string `json:"token,omitempty"`
	NewBalance *Rates `json:"new_balance,omitempty"`
}

type AccountChange struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type exchange struct {
	From   string  `json:"from_currency"`
	To     string  `json:"to_currency"`
	Amount float64 `json:"amount,omitempty"`
}
