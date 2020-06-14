package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Student struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	RollNo string `gorm:"size:255;unique;not null" json:"roll_no"`
	StudentName string `gorm:"size:255;not null" json:"student_name"`
	Major string `gorm:"size:255;not null" json:"major"`
	Year string `gorm:"size:255;not null" json:"year"`
	Email string `gorm:"size:255;not null" json:"email"`
	Phone string `gorm:"size:255;not null" json:"phone"`
	Address string `gorm:"size:255;not null" json:"address"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (student *Student) Prepare() {
	student.ID = 0
	student.RollNo = html.EscapeString(strings.TrimSpace(student.RollNo))
	student.StudentName = html.EscapeString(student.StudentName)
	student.Major = html.EscapeString(strings.TrimSpace(student.Major))
	student.Year = html.EscapeString(student.Year)
	student.Email = html.EscapeString(strings.TrimSpace(student.Email))
	student.Phone = html.EscapeString(student.Phone)
	student.Address = html.EscapeString(student.Address)
}

func (student *Student) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if student.RollNo == "" {
			return errors.New("Required RollNo ")
		}
		if student.StudentName == "" {
			return errors.New("Required Student Name ")
		}
		if student.Major == "" {
			return errors.New("Required Major ")
		}
		if student.Year == "" {
			return errors.New("Required Year ")
		}
		if student.Email == "" {
			return errors.New("Required Email ")
		}
		if err := checkmail.ValidateFormat(student.Email); err != nil {
			return errors.New("Invalid Email ")
		}
		if student.Phone == "" {
			return errors.New("Required Phone ")
		}
		if student.Address == "" {
			return errors.New("Required Address ")
		}
		return nil

	case "login":
		if student.RollNo == "" {
			return errors.New("Required RollNo ")
		}
		return nil

	default:
		if student.RollNo == "" {
			return errors.New("Required RollNo ")
		}
		if student.StudentName == "" {
			return errors.New("Required Student Name ")
		}
		if student.Major == "" {
			return errors.New("Required Major ")
		}
		if student.Year == "" {
			return errors.New("Required Year ")
		}
		if student.Email == "" {
			return errors.New("Required Email ")
		}
		if err := checkmail.ValidateFormat(student.Email); err != nil {
			return errors.New("Invalid Email ")
		}
		if student.Phone == "" {
			return errors.New("Required Phone ")
		}
		if student.Address == "" {
			return errors.New("Required Address ")
		}
		return nil
	}
}

func (student *Student) SaveStudent(db *gorm.DB) (*Student, error) {
	var err error
	err = db.Debug().Create(&student).Error
	if err != nil {
		return &Student{}, err
	}
	return student, err
}

func (student *Student) FindAllStudents(db *gorm.DB) (*[]Student, error) {
	var err error
	students := []Student{}
	err = db.Debug().Model(&Student{}).Limit(100).Find(&students).Error
	if err != nil {
		return &[]Student{}, err
	}
	return &students, err
}

func (student *Student) FindStudentByID(db *gorm.DB, sid uint32) (*Student, error) {
	var err error
	err = db.Debug().Model(&Student{}).Where("id = ?", sid).Take(&student).Error
	if err != nil {
		return &Student{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Student{}, errors.New("Student not found ")
	}
	return student, err
}

func (student *Student) UpdateStudent(db *gorm.DB, sid uint32) (*Student, error) {
	db = db.Debug().Model(&Student{}).Where("id = ?", sid).Take(&Student{}).UpdateColumns(
		map[string]interface{}{
			"roll_no":      student.RollNo,
			"student_name": student.StudentName,
			"major":        student.Major,
			"year":         student.Year,
			"email":        student.Email,
			"phone":        student.Phone,
			"address":      student.Address,
		},
	)
	if db.Error != nil {
		return &Student{}, nil
	}
	err := db.Debug().Model(&Student{}).Where("id = ?", sid).Take(&student).Error
	if err != nil {
		return &Student{}, err
	}
	return student, err
}

func (student *Student) DeleteStudent(db *gorm.DB, sid uint32) (int64, error) {
	db = db.Debug().Model(&Student{}).Where("id = ?", sid).Take(&Student{}).Delete(&Student{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
