package auth

import (
	"encoding/json"
	"fmt"
	"http-server/configs"
	"http-server/pkg/res"
	"net/http"

	"github.com/go-playground/validator/v10"
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

		validate := validator.New()
		err = validate.Struct(payload)

		if err != nil {
			res.Json(w, err.Error(), 402)
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
