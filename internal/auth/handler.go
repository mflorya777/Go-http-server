package auth

import (
	"encoding/json"
	"fmt"
	"http-server/configs"
	"http-server/pkg/res"
	"net/http"
	"regexp"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}

	router.HandleFunc("/auth/login", handler.Login())
	router.HandleFunc("/auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Login")

		// Прочитать Body
		var payload LoginRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			res.Json(w, err.Error(), 402)
			return
		}
		if payload.Email == "" {
			res.Json(w, "Email required", 402)
			return
		}

		// Валидация Email с regexp
		match, _ := regexp.MatchString(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`, payload.Email)
		if !match {
			res.Json(w, "Wrong Email", 402)
			return
		}

		if payload.Password == "" {
			res.Json(w, "Password required", 402)
			return
		}
		fmt.Println(payload)

		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, 200)
	}

}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}

}
