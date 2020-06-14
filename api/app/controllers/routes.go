package controllers

import "github.com/OoThan/students-attendance-go-api/api/app/middlewares"

func (server *Server) initializeRoutes() {

	//Home Route
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")

	//login faculty to CURD teacher
	server.Router.HandleFunc("/faculty/login", middlewares.SetMiddlewareJSON(server.FacultyLogin)).Methods("POST")
	server.Router.HandleFunc("/student/login", middlewares.SetMiddlewareJSON(server.StudentLogin)).Methods("POST")
	server.Router.HandleFunc("/teacher/login", middlewares.SetMiddlewareJSON(server.TeacherLogin)).Methods("POST")

	//Faculty routes
	server.Router.HandleFunc("/faculties", middlewares.SetMiddlewareJSON(server.CreateFaculty)).Methods("POST")
	server.Router.HandleFunc("/faculties", middlewares.SetMiddlewareJSON(server.GetFaculties)).Methods("GET")
	server.Router.HandleFunc("/faculty/{id}", middlewares.SetMiddlewareJSON(server.GetFaculty)).Methods("GET")
	server.Router.HandleFunc("/faculty/{id}", middlewares.SetMiddlewareJSON(server.UpdateFaculty)).Methods("PUT")
	server.Router.HandleFunc("/faculty/{id}", middlewares.SetMiddlewareJSON(server.DeleteFaculty)).Methods("DELETE")

	//Teacher routes
	server.Router.HandleFunc("/teachers", middlewares.SetMiddlewareJSON(server.GetTeachers)).Methods("GET")
	server.Router.HandleFunc("/teacher/{id}", middlewares.SetMiddlewareJSON(server.GetTeacher)).Methods("GET")
	//teacher CUD with faculty token
	server.Router.HandleFunc("/teachers", middlewares.SetMiddlewareJSON(server.CreateTeacher)).Methods("POST")
	server.Router.HandleFunc("/teacher/faculty/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateTeacher))).Methods("PUT")
	server.Router.HandleFunc("/teacher/faculty/{id}", middlewares.SetMiddlewareAuthentication(server.DeleteTeacher)).Methods("DELETE")

	//Student routes
	server.Router.HandleFunc("/students", middlewares.SetMiddlewareJSON(server.CreateStudent)).Methods("POST")
	server.Router.HandleFunc("/students", middlewares.SetMiddlewareJSON(server.GetStudents)).Methods("GET")
	server.Router.HandleFunc("/student/{id}", middlewares.SetMiddlewareJSON(server.GetStudent)).Methods("GET")
	server.Router.HandleFunc("/student/{id}", middlewares.SetMiddlewareJSON(server.UpdateStudent)).Methods("PUT")
	server.Router.HandleFunc("/student/{id}", middlewares.SetMiddlewareJSON(server.DeleteStudent)).Methods("DELETE")

	//Attendance are under control of Teacher for CURD
	server.Router.HandleFunc("/teacher/attendance/{roll_no}", middlewares.SetMiddlewareJSON(server.CreateAttendance)).Methods("POST")
	server.Router.HandleFunc("/teacher/attendances", middlewares.SetMiddlewareJSON(server.GetAllAttendances)).Methods("GET")
	server.Router.HandleFunc("/teacher/attendance/{id}", middlewares.SetMiddlewareJSON(server.GetAttendanceByID)).Methods("GET")
	server.Router.HandleFunc("/teacher/present-attendance/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.PresentStudent))).Methods("PUT")
	server.Router.HandleFunc("/teacher/absent-attendance/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.AbsentStudent))).Methods("PUT")
	server.Router.HandleFunc("/teacher/delete-attendance/{id}", middlewares.SetMiddlewareAuthentication(server.DeleteAttendance)).Methods("DELETE")
}
