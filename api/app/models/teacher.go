package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Teacher struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	TeacherName string `gorm:"size:255;not null" json:"teacher_name"`
	Course string `gorm:"size:255;not null" json:"course"`
	Email string `gorm:"size:255;not null" json:"email"`
	Phone string `gorm:"size:255;not null" json:"phone"`
	Address string `gorm:"size:255;not null" json:"address"`
	Faculty Faculty `json:"faculty"`
	FacultyID uint32 `gorm:"not null" json:"faculty_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (teacher *Teacher) Prepare() {
	teacher.ID = 0
	teacher.TeacherName = html.EscapeString(teacher.TeacherName)
	teacher.Course = html.EscapeString(teacher.Course)
	teacher.Email = html.EscapeString(strings.TrimSpace(teacher.Email))
	teacher.Phone = html.EscapeString(strings.TrimSpace(teacher.Phone))
	teacher.Address = html.EscapeString(teacher.Address)
	teacher.Faculty = Faculty{}
	teacher.CreatedAt = time.Now()
	teacher.UpdatedAt = time.Now()
}

func (teacher *Teacher) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if teacher.TeacherName == "" {
			return errors.New("Required Teacher Name ")
		}
		if teacher.Course == "" {
			return errors.New("Required Course ")
		}
		if teacher.Email == "" {
			return errors.New("Required Email ")
		}
		if err := checkmail.ValidateFormat(teacher.Email); err != nil {
			return errors.New("Invalid Email ")
		}
		if teacher.Phone == "" {
			return errors.New("Required Phone ")
		}
		if teacher.Address == "" {
			return errors.New("Required Address ")
		}
		if teacher.FacultyID < 1 {
			return errors.New("Required Faculty ")
		}
		return nil

	case "login":
		if teacher.Email == "" {
			return errors.New("Required RollNo ")
		}
		if err := checkmail.ValidateFormat(teacher.Email); err != nil {
			return errors.New("Invalid Email ")
		}
		return nil

	default:
		if teacher.TeacherName == "" {
			return errors.New("Required Teacher Name ")
		}
		if teacher.Course == "" {
			return errors.New("Required Course ")
		}
		if teacher.Email == "" {
			return errors.New("Required Email ")
		}
		if err := checkmail.ValidateFormat(teacher.Email); err != nil {
			return errors.New("Invalid Email ")
		}
		if teacher.Phone == "" {
			return errors.New("Required Phone ")
		}
		if teacher.Address == "" {
			return errors.New("Required Address ")
		}
		if teacher.FacultyID < 1 {
			return errors.New("Required Faculty ")
		}
		return nil
	}
}

func (teacher *Teacher) SaveTeacher(db *gorm.DB) (*Teacher, error) {
	var err error
	err = db.Debug().Model(&Teacher{}).Create(&teacher).Error
	if err != nil {
		return &Teacher{}, err
	}
	if teacher.ID != 0 {
		err = db.Debug().Model(&Faculty{}).Where("id = ?", teacher.FacultyID).Error
		if err != nil {
			return &Teacher{}, err
		}
	}
	return teacher, nil
}

func (teacher *Teacher) FindAllTeachers(db *gorm.DB) (*[]Teacher, error) {
	var err error
	teachers := []Teacher{}
	err = db.Debug().Model(&Teacher{}).Limit(100).Find(&teachers).Error
	if err != nil {
		return &[]Teacher{}, err
	}
	if len(teachers) > 0 {
		for i, _ := range teachers {
			err := db.Debug().Model(&Faculty{}).Where("id = ?", teachers[i].FacultyID).Take(&teachers[i].Faculty).Error
			if err != nil {
				return &[]Teacher{}, err
			}
		}
	}
	return &teachers, nil
}

func (teacher *Teacher) FindTeacherByID(db *gorm.DB, tid uint32) (*Teacher, error) {
	var err error
	err = db.Debug().Model(&Teacher{}).Where("id = ?", tid).Take(&teacher).Error
	if err != nil {
		return &Teacher{}, err
	}
	if teacher.ID != 0 {
		err = db.Debug().Model(&Faculty{}).Where("id = ?", teacher.FacultyID).Take(&teacher.Faculty).Error
		if err != nil {
			return &Teacher{}, err
		}
	}
	return teacher, nil
}

func (teacher *Teacher) UpdateTeacher(db *gorm.DB) (*Teacher, error) {
	var err error
	err = db.Debug().Model(&Teacher{}).Where("id = ?", teacher.ID).UpdateColumn(
		Teacher{
			TeacherName: teacher.TeacherName,
			Course: teacher.Course,
			Email: teacher.Email,
			Phone: teacher.Phone,
			Address: teacher.Address,
			UpdatedAt: time.Now(),
		}).Error
	if err != nil {
		return &Teacher{}, err
	}
	if teacher.ID != 0 {
		err = db.Debug().Model(&Faculty{}).Where("id = ?", teacher.FacultyID).Take(&teacher.Faculty).Error
		if err != nil {
			return &Teacher{}, err
		}
	}
	return teacher, nil
}

func (teacher *Teacher) DeleteTeacher(db *gorm.DB, tid, fid uint32) (int64, error) {
	db = db.Debug().Model(&Teacher{}).Where("id = ? and faculty_id = ?", tid, fid).Take(&Teacher{}).Delete(&Teacher{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Teacher not Found ")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
