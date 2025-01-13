package handlers

import "wallet/internal/service"

func NewHandler(s *service.Service) *Handler {
	return &Handler{Service: s}
}
