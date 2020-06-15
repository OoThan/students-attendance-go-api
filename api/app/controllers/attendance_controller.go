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

func (server *Server) CreateAttendance(w http.ResponseWriter, r *http.Request) {
	//to get student roll_no from URL
	vars := mux.Vars(r)
	student_roll_no := vars["roll_no"]
	fmt.Println(student_roll_no)

	//to get teacher id from token
	tid, err := auth.ExtractTeacherTokenID(r)
	if err != nil {
		fmt.Println(tid)
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized "))
		return
	}

	//to get data from body body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("body error"))
		return
	}
	attendance := models.Attendance{}
	err = json.Unmarshal(body, &attendance)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("json unmarshal"))
		return
	}
	attendance.TeacherID = tid
	attendance.StudentRollNo = student_roll_no
	attendanceCreated, err := attendance.SaveAttendance(server.DB, student_roll_no)
	//fmt.Println(attendanceCreated)
	if err != nil {
		formattedError := formaterrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	//if attendance.TeacherID == attendanceCreated.TeacherID {
	//	err = server.DB.Debug().Model(&models.Teacher{}).Where("id = ?", attendanceCreated.TeacherID).Take(&attendanceCreated.Teacher).Error
	//	if err != nil {
	//		responses.ERROR(w, http.StatusUnauthorized, errors.New("attendance.TeacherID"))
	//		return
	//	}
	//}
	if attendance.StudentRollNo == attendanceCreated.StudentRollNo {
		fmt.Println(attendance.StudentRollNo, attendanceCreated.StudentRollNo)
		err = server.DB.Debug().Model(&models.Student{}).Where("roll_no = ?", attendanceCreated.StudentRollNo).Take(&attendanceCreated.Student).Error
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("roll_no error"))
			fmt.Println(attendance.StudentRollNo, attendanceCreated.StudentRollNo)
			return
		}
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, attendanceCreated.ID))
	responses.JSON(w, http.StatusCreated, attendanceCreated)
}

func (server *Server) GetAllAttendances(w http.ResponseWriter, r *http.Request) {
	attendance := models.Attendance{}
	attendances, err := attendance.FindAllAttendances(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, attendances)
}

func (server *Server) GetAttendanceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	attendance := models.Attendance{}
	attendanceGotten, err := attendance.FindAttendanceByID(server.DB, uint32(aid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, attendanceGotten)
}

func (server *Server) AbsentStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tid, err := auth.ExtractTeacherTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized "))
		return
	}
	attendance := models.Attendance{}
	err = server.DB.Debug().Model(&models.Attendance{}).Where("id = ?", aid).Take(&attendance).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Record Not Found! "))
		return
	}
	if tid != attendance.TeacherID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized "))
		return
	}
	attendanceUpdated, err := attendance.SetAbsent(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, attendanceUpdated)
}

func (server *Server) PresentStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tid, err := auth.ExtractTeacherTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized "))
		return
	}
	attendance := models.Attendance{}
	err = server.DB.Debug().Model(&models.Attendance{}).Where("id = ?", aid).Take(&attendance).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Record Not Found "))
		return
	}
	if tid != attendance.TeacherID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized "))
		return
	}
	attendanceUpdated, err := attendance.SetPresent(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, attendanceUpdated)
}

func (server *Server) DeleteAttendance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tid, err := auth.ExtractTeacherTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized "))
		return
	}
	attendance := models.Attendance{}
	err = server.DB.Debug().Model(&models.Attendance{}).Where("id = ?", aid).Take(&attendance).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Record Not Found! "))
		return
	}
	if tid != attendance.TeacherID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized "))
		return
	}
	attendanceDeleted, err := attendance.DeleteAttendance(server.DB, uint32(aid), tid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	err = server.DB.Model(&models.Faculty{}).Where("id = ?", attendanceDeleted.Teacher.FacultyID).Take(&attendanceDeleted.Teacher.Faculty).Error
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized "))
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", aid))
	responses.JSON(w, http.StatusOK, attendanceDeleted)
}
