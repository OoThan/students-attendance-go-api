package controllers

import (
	"github.com/OoThan/students-attendance-go-api/api/app/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To My Student Attendance Management Awesome API ><")
}
