package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OoThan/students-attendance-go-api/api/app/auth"
	"github.com/OoThan/students-attendance-go-api/api/app/models"
	"github.com/OoThan/students-attendance-go-api/api/app/responses"
	"github.com/OoThan/students-attendance-go-api/api/app/utils/formaterrors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	teacher := models.Teacher{}
	err = json.Unmarshal(body, &teacher)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	teacher.Prepare()
	err = teacher.Validate("")
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}
	fid, err := auth.ExtractFacultyTokenID(r)
	fmt.Println(fid, err)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized fid "))
		return
	}
	//if fid != teacher.FacultyID {
	//	fmt.Println(fid, teacher.FacultyID)
	//	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	//	return
	//}
	teacher.FacultyID = fid
	teacherCreated, err := teacher.SaveTeacher(server.DB)
	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	if teacher.FacultyID == teacherCreated.FacultyID {
		err = server.DB.Debug().Model(&models.Faculty{}).Where("id = ?", teacherCreated.FacultyID).Take(&teacherCreated.Faculty).Error
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		}
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, teacherCreated.ID))
	responses.JSON(w, http.StatusCreated, teacherCreated)
}

func (server *Server) GetTeachers(w http.ResponseWriter, r *http.Request) {
	teacher := models.Teacher{}
	teachers, err := teacher.FindAllTeachers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, teachers)
}

func (server *Server) GetTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	teacher := models.Teacher{}
	teacherReceived, err := teacher.FindTeacherByID(server.DB, uint32(tid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, teacherReceived)
}

func (server *Server) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//Check if the teacher id ia valid
	tid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//check if the faculty id ia vaild and get the teacher id form it
	fid, err := auth.ExtractFacultyTokenID(r)
	if err != nil {
		fmt.Println(fid)
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized tokenID "))
		return
	}

	//check if the teacher exist
	teacher := models.Teacher{}
	err = server.DB.Debug().Model(models.Teacher{}).Where("id = ?", tid).Take(&teacher).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Post not found "))
		return
	}

	//if a faculty attempt to update a teacher not belonging to him
	if fid != teacher.FacultyID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized fid "))
		return
	}
	//Read the data teacherUpdate
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	//start processing the request data
	//teacherUpdate := models.Teacher{}
	err = json.Unmarshal(body, &teacher)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if fid != teacher.FacultyID {
		fmt.Println(fid, teacher.FacultyID)
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized teacherUpdate.FacultyID "))
		return
	}
	teacher.Prepare()
	err = teacher.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//teacherUpdate.ID = teacher.ID
	teacherUpdated, err := teacher.UpdateTeacher(server.DB)
	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, teacherUpdated)
}

func (server *Server) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fid, err := auth.ExtractFacultyTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized "))
		return
	}
	teacher := models.Teacher{}
	err = server.DB.Debug().Model(&models.Teacher{}).Where("id = ?", tid).Take(&teacher).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized "))
		return
	}
	if fid != teacher.FacultyID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized "))
		return
	}
	_, err = teacher.DeleteTeacher(server.DB, uint32(tid), fid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", tid))
	responses.JSON(w, http.StatusNoContent, "")
}
