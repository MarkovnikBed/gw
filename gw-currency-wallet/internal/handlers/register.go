package handlers

import (
	"encoding/json"
	"net/http"
)

// @Summary      Регистрация
// @Description  Регистрация нового пользователя
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user_data  body  UserData  true  "Данные пользователя"
// @Success      201 {object} response "пользователь успешно создан"
// @Failure      400 {object} response "некорректный запрос"
// @Failure      409 {object} response "пользователь с таким именем или почтой уже существует"
// @Router       /api/v1/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	ud := UserData{}

	if err := json.NewDecoder(r.Body).Decode(&ud); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&response{Error: "некорректный запрос"})
		return
	}

	if err := h.Service.Repository.Registration(ud.Username, ud.Email, ud.Password); err != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&response{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&response{Message: "пользователь успешно создан"})

}
