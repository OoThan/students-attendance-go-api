package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/OoThan/students-attendance-go-api/api/app/models"
	"github.com/OoThan/students-attendance-go-api/api/app/responses"
	"github.com/OoThan/students-attendance-go-api/api/app/utils/formaterrors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateStudent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	student := models.Student{}
	err = json.Unmarshal(body, &student)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	student.Prepare()
	err = student.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	studentCreated, err := student.SaveStudent(server.DB)
	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, studentCreated.ID))
	responses.JSON(w, http.StatusCreated, studentCreated)
}

func (server *Server) GetStudents(w http.ResponseWriter, r *http.Request) {
	student := models.Student{}
	students, err := student.FindAllStudents(server.DB)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, students)
}

func (server *Server) GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	student := models.Student{}
	studentGotten, err := student.FindStudentByID(server.DB, uint32(sid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, studentGotten)
}

func (server *Server) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["id"], 10 ,32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	student := models.Student{}
	err = json.Unmarshal(body, &student)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	student.Prepare()
	err = student.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedStudent, err := student.UpdateStudent(server.DB, uint32(sid))
	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedStudent)
}

func (server *Server) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	student := models.Student{}
	sid, err := strconv.ParseUint(vars["id"], 10 ,32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = student.DeleteStudent(server.DB, uint32(sid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", sid))
	responses.JSON(w, http.StatusNoContent, "")
}
