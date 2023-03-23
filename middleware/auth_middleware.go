package middleware

import (
	"github.com/ariopri/golang_restfull_api/helper"
	"github.com/ariopri/golang_restfull_api/models/web"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	err := godotenv.Load()
	helper.PanicIfError(err)
	appKey := os.Getenv("APP_KEY")

	if appKey == r.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helper.WriteToResponseBody(w, webResponse)
	}
}
