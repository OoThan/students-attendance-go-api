package seed

import (
	"github.com/OoThan/students-attendance-go-api/api/app/models"
	"github.com/jinzhu/gorm"
	"log"
)

var faculties = []models.Faculty{
	models.Faculty{
		Acronym: "FCS",
		FacultyName: "Faculty of Computer Science",
		FacultyMail: "fcs@ucssittway.edu.mm",
	},
	models.Faculty{
		Acronym: "FCT",
		FacultyName: "Faculty of Computer Technology",
		FacultyMail: "fct@ucssittway.edu.mm",
	},
	models.Faculty{
		Acronym: "FCA",
		FacultyName: "Faculty of Computer Application",
		FacultyMail: "fca@ucssittway.edu.mm",
	},
}

var teachers = []models.Teacher{
	models.Teacher{
		TeacherName: "Dr. Mya Thi",
		Course: "Java",
		Email: "myathi@ucssittway.edu.mm",
		Phone: "09421712345",
		Address: "Sittway",
	},
	models.Teacher{
		TeacherName: "Dr. Thein Maung",
		Course: "Embedded System",
		Email: "theinmaung@ucssittway.edu.mm",
		Phone: "09262963456",
		Address: "Minbya",
	},
	models.Teacher{
		TeacherName: "Dr. Mya Shwe",
		Course: "Accounting",
		Email: "myashwe@ucssittway.edu.mm",
		Address: "Yangon",
	},
}

var students = []models.Student{
	models.Student{
		RollNo: "5CS-1",
		StudentName: "Maung Oo Kyaw Than",
		Major: "CS",
		Year: "Fifth year",
		Email: "maungoothan@ucssittway.edu.mm",
		Phone: "09250014910",
		Address: "Minbya, Rakhine State",
	},
	models.Student{
		RollNo: "5CT-1",
		StudentName: "Maung Myint Thein",
		Major: "CT",
		Year: "Fifth year",
		Email: "myintthein@ucssittway.edu.mm",
		Phone: "09780567432",
		Address: "Minbya, Rakhine state",
	},
	models.Student{
		RollNo: "5CS-2",
		StudentName: "Ma Nandar Khine",
		Major: "CS",
		Year: "Fifth year",
		Email: "nandarkhine@ucssittway.edu.mm",
		Phone: "09260786543",
		Address: "Sittway, Rakhine State",
	},
}

/*var attendances = []models.Attendance{
	models.Attendance{
		Status: true,
	},
	models.Attendance{
		Status: false,
	},
	models.Attendance{
		Status: true,
	},
}*/

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Attendance{}, &models.Teacher{}, &models.Faculty{}, &models.Student{}).Error
	if err != nil {
		log.Fatalf("cannot drop tables: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Faculty{}, &models.Teacher{}, &models.Student{}, &models.Attendance{}).Error
	if err != nil {
		log.Fatalf("cannot migrate tables: %v", err)
	}
	err = db.Debug().Model(&models.Teacher{}).AddForeignKey("faculty_id", "faculties(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching forigen key error: %v", err)
	}
	//Insert data for faculty and teacher
	for i, _ := range faculties {
		err = db.Debug().Model(&models.Faculty{}).Create(&faculties[i]).Error
		if err != nil {
			log.Fatalf("cannot seed faulties table: %v", err)
		}
		teachers[i].FacultyID = faculties[i].ID
		err = db.Debug().Model(&models.Teacher{}).Create(&teachers[i]).Error
		if err != nil {
			log.Fatalf("cannot seed teachers table: %v", err)
		}
	}
	//Insert data for student
	for i, _ := range students {
		err = db.Debug().Model(&models.Student{}).Create(&students[i]).Error
		if err != nil {
			log.Fatalf("cannot seed students table: %v", err)
		}
	}
	//Insert data to attendance
	err = db.Debug().Model(&models.Attendance{}).AddForeignKey("teacher_id", "teachers(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key errpr: %v", err)
	}
	/*for i, _ := range teachers {
		attendances[i].TeacherID = teachers[i].ID
		err = db.Debug().Model(&models.Attendance{}).Create(&attendances[i]).Error
		if err != nil {
			log.Fatalf("cannot seed attendances table: %v", err)
		}
	}*/
}
