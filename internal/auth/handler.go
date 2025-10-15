package auth

import (
	"fmt"
	"http-server/configs"
	"http-server/pkg/req"
	"http-server/pkg/res"
	"net/http"
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

		body, err := req.HandleBody[LoginRequest](&w, r)

		if err != nil {
			return
		}

		fmt.Println(body)

		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, 200)
	}

}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)

		if err != nil {
			return
		}

		fmt.Println(body)

	}

}
