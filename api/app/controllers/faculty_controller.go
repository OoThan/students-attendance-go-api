package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OoThan/students-attendance-go-api/api/app/models"
	"github.com/OoThan/students-attendance-go-api/api/app/responses"
	"github.com/OoThan/students-attendance-go-api/api/app/utils/formaterrors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateFaculty(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	faculty := models.Faculty{}
	err = json.Unmarshal(body, &faculty)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	faculty.Prepare()
	err = faculty.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	facultyCreated, err := faculty.SaveFaculty(server.DB)
	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, facultyCreated.ID))
	responses.JSON(w, http.StatusCreated, facultyCreated)
}

func (server *Server) GetFaculties(w http.ResponseWriter, r *http.Request) {
	faculty := models.Faculty{}
	faculties, err := faculty.FindAllFaculties(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, faculties)
}

func (server *Server) GetFaculty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	faculty := models.Faculty{}
	facultyGotten, err := faculty.FindFacultyByID(server.DB, uint32(fid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, facultyGotten)
}

func (server *Server) UpdateFaculty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	faculty := models.Faculty{}
	err = json.Unmarshal(body, &faculty)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	faculty.Prepare()
	err = faculty.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedFaculty, err := faculty.UpdateFaculty(server.DB, uint32(fid))
	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	if updatedFaculty.ID == 0 {
		responses.ERROR(w, http.StatusNotFound, errors.New("Faculty not found "))
		return
	}
	responses.JSON(w, http.StatusOK, updatedFaculty)
}

func (server *Server) DeleteFaculty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	faculty := models.Faculty{}
	fid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = faculty.DeleteFaculty(server.DB, uint32(fid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", fid))
	responses.JSON(w, http.StatusNoContent, "")
}
