package middlewares

import (
	"errors"
	"github.com/OoThan/students-attendance-go-api/api/app/auth"
	"github.com/OoThan/students-attendance-go-api/api/app/responses"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		next(writer, request)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := auth.TokenValid(request)
		if err != nil {
			responses.ERROR(writer, http.StatusUnauthorized, errors.New("Unauthorized "))
			return
		}
		next(writer, request)
	}
}
