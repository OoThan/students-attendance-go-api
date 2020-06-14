package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/OoThan/students-attendance-go-api/api/app/auth"
	"github.com/OoThan/students-attendance-go-api/api/app/models"
	"github.com/OoThan/students-attendance-go-api/api/app/responses"
	"io/ioutil"
	"net/http"
)

func (server *Server) FacultyLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w ,http.StatusUnprocessableEntity, err)
		return
	}
	faculty := models.Faculty{}
	err = json.Unmarshal(body, &faculty)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	faculty.Prepare()
	err = faculty.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.FacultySignIn(faculty.Acronym)
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) FacultySignIn(acronym string) (string, error) {
	var err error
	faculty := models.Faculty{}
	err = server.DB.Debug().Model(models.Faculty{}).Where("acronym = ?", acronym).Take(&faculty).Error
	if err != nil {
		return "", err
	}
	fmt.Println(acronym, faculty.ID)
	return auth.CreateFacultyToken(faculty.ID)
}

func (server *Server) StudentLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w ,http.StatusUnprocessableEntity, err)
		return
	}
	student := models.Student{}
	err = json.Unmarshal(body, &student)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	student.Prepare()
	err = student.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.StudentSignIn(student.RollNo)
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) StudentSignIn(roll_no string) (string, error) {
	var err error
	student := models.Student{}
	err = server.DB.Debug().Model(models.Student{}).Where("roll_no = ?", roll_no).Take(&student).Error
	if err != nil {
		return "", err
	}
	fmt.Println(roll_no, student.ID)
	return auth.CreateStudentToken(student.ID)
}

func (server *Server) TeacherLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w ,http.StatusUnprocessableEntity, err)
		return
	}
	teacher := models.Teacher{}
	err = json.Unmarshal(body, &teacher)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	teacher.Prepare()
	err = teacher.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.TeacherSignIn(teacher.Email)
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) TeacherSignIn(email string) (string, error) {
	var err error
	teacher := models.Teacher{}
	err = server.DB.Debug().Model(models.Teacher{}).Where("email = ?", email).Take(&teacher).Error
	if err != nil {
		return "", err
	}
	fmt.Println(email, teacher.ID)
	return auth.CreateTeacherToken(teacher.ID)
}


