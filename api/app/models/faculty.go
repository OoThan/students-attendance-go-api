package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Faculty struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Acronym string `gorm:"size:255;not null;unique" json:"acronym"`
	FacultyName string `gorm:"size:255;not null" json:"faculty_name"`
	FacultyMail string `gorm:"size:255;not null" json:"faculty_mail"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (faculty *Faculty) Prepare() {
	faculty.ID = 0
	faculty.Acronym = html.EscapeString(strings.TrimSpace(faculty.Acronym))
	faculty.FacultyName = html.EscapeString(faculty.FacultyName)
	faculty.FacultyMail = html.EscapeString(strings.TrimSpace(faculty.FacultyMail))
	faculty.CreatedAt = time.Now()
	faculty.UpdatedAt = time.Now()
}

func (faculty *Faculty) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if faculty.Acronym == "" {
			return errors.New("Required Faculty Acronym ")
		}
		if faculty.FacultyName == "" {
			return errors.New("Required Faculty Name ")
		}
		if faculty.FacultyMail == "" {
			return errors.New("Required Faculty Mail ")
		}
		if err := checkmail.ValidateFormat(faculty.FacultyMail); err != nil {
			return errors.New("Invalid Email ")
		}
		return nil
	case "login":
		if faculty.Acronym == "" {
			return errors.New("Required acronym ")
		}
		return nil
	default:
		if faculty.Acronym == "" {
			return errors.New("Required Faculty Acronym ")
		}
		if faculty.FacultyName == "" {
			return errors.New("Required Faculty Name ")
		}
		if faculty.FacultyMail == "" {
			return errors.New("Required Faculty Mail ")
		}
		if err := checkmail.ValidateFormat(faculty.FacultyMail); err != nil {
			return errors.New("Invalid Email ")
		}
		return nil
	}
}

func (faculty *Faculty) SaveFaculty(db *gorm.DB) (*Faculty, error) {
	var err error
	err = db.Debug().Create(&faculty).Error
	if err != nil {
		return &Faculty{}, err
	}
	return faculty, nil
}

func (faculty *Faculty) FindAllFaculties(db *gorm.DB) (*[]Faculty, error) {
	var err error
	faculties := []Faculty{}
	err = db.Debug().Model(&Faculty{}).Limit(100).Find(&faculties).Error
	if err != nil {
		return &[]Faculty{}, err
	}
	return &faculties, nil
}

func (faculty *Faculty) FindFacultyByID(db *gorm.DB, fid uint32) (*Faculty, error) {
	var err error
	err = db.Debug().Model(Faculty{}).Where("id = ?", fid).Take(&faculty).Error
	if err != nil {
		return &Faculty{}, err
	}
	return faculty, nil
}

func (faculty *Faculty) UpdateFaculty(db *gorm.DB, fid uint32) (*Faculty, error) {
	db = db.Debug().Model(&Faculty{}).Where("id = ?", fid).Take(&Faculty{}).UpdateColumn(
		map[string]interface{}{
		"acronym": faculty.Acronym,
		"faculty_name": faculty.FacultyName,
		"faculty_mail": faculty.FacultyMail,
		"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Faculty{}, db.Error
	}
	err := db.Debug().Model(&Faculty{}).Where("id = ?", fid).Take(&faculty).Error
	if err != nil {
		return &Faculty{}, err
	}
	return faculty, nil
}

func (faculty *Faculty) DeleteFaculty(db *gorm.DB, fid uint32) (int64, error) {
	db = db.Debug().Model(&Faculty{}).Where("id = ?", fid).Take(&Faculty{}).Delete(&Faculty{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}


