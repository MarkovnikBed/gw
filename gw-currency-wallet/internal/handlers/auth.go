package handlers

import "net/http"

func (h *Handler) Auth(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "пустой заголовок авторизации", http.StatusUnauthorized)
			return
		}
		if h.Service.Repository.Auth(token) {
			handler(w, r)
		} else {
			http.Error(w, "не удалось провести авторизацию", http.StatusUnauthorized)
		}
	}
}
