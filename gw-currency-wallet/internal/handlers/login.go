package handlers

import (
	"encoding/json"
	"net/http"
)

// @Summary      Авторизация
// @Description  Авторизация пользователя
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user_data  body  UserData  true  "Данные пользователя"
// @Success      200 {object} response "токен авторизации"
// @Failure      401 {object} response "неверный логин или пароль"
// @Router       /api/v1/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	defer r.Body.Close()

	ud := UserData{}
	json.NewDecoder(r.Body).Decode(&ud)

	token, err := h.Service.Repository.Login(ud.Username, ud.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response{Error: err.Error()})

		return
	}
	json.NewEncoder(w).Encode(&response{Token: token})

}
