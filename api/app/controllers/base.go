package controllers

import (
	"fmt"
	"github.com/OoThan/students-attendance-go-api/api/app/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DBDriver, DBUser, DBPassword, DBPort, DBHost, DBName string) {
	var err error
	DbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	server.DB, err = gorm.Open(DBDriver, DbURL)
	if err != nil {
		fmt.Printf("cannot connect to %s database ", DBDriver)
		log.Fatal("This is error: ", err)
	}
	server.DB.Debug().AutoMigrate(&models.Faculty{}, &models.Teacher{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
