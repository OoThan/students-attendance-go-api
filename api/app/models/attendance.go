package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Attendance struct {
	ID            uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Status        bool      `gorm:"not null" json:"status"`
	TeacherID     uint32    `gorm:"not null" json:"teacher_id"`
	Teacher       Teacher   `json:"teacher"`
	StudentRollNo string    `gorm:"not null" json:"student_roll_no"`
	Student       Student   `json:"student"`
	UpdateAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
	DeleteAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"delete_at"`
}

func (attendance *Attendance) SaveAttendance(db *gorm.DB, student_roll_no string) (*Attendance, error) {
	var err error
	err = db.Debug().Model(&Attendance{}).Create(&attendance).Error
	if err != nil {
		return &Attendance{}, err
	}
	if attendance.ID != 0 {
		err = db.Debug().Model(&Teacher{}).Where("id = ?", attendance.TeacherID).Take(&attendance.Teacher).Error
		if err != nil {
			return &Attendance{}, err
		}
		err = db.Debug().Model(&Student{}).Where("roll_no = ?", student_roll_no).Take(&attendance.Student).Error
		if err != nil {
			return &Attendance{}, err
		}
		err = db.Debug().Model(&Faculty{}).Where("id = ?", attendance.Teacher.FacultyID).Take(&attendance.Teacher.Faculty).Error
		if err != nil {
			return &Attendance{}, err
		}
	}
	return attendance, nil
}

func (attendance *Attendance) FindAllAttendances(db *gorm.DB) (*[]Attendance, error) {
	var err error
	attendances := []Attendance{}
	err = db.Debug().Model(&Attendance{}).Limit(100).Find(&attendances).Error
	fmt.Println(attendances, len(attendances))
	if err != nil {
		return &[]Attendance{}, err
	}
	if len(attendances) > 0 {
		for i, _ := range attendances {
			err = db.Debug().Model(&Teacher{}).Where("id = ?", uint32(attendances[i].TeacherID)).Take(&attendances[i].Teacher).Error
			if err != nil {
				return &[]Attendance{}, err
			}
			err = db.Debug().Model(&Student{}).Where("roll_no = ?", attendances[i].StudentRollNo).Take(&attendances[i].Student).Error
			if err != nil {
				return &[]Attendance{}, err
			}
			err = db.Debug().Model(&Faculty{}).Where("id = ?", attendances[i].Teacher.FacultyID).Take(&attendances[i].Teacher.Faculty).Error
			if err != nil {
				return &[]Attendance{}, err
			}
		}
	}
	return &attendances, nil
}

func (attendance *Attendance) FindAttendanceByID(db *gorm.DB, aid uint32) (*Attendance, error) {
	var err error
	err = db.Debug().Model(&Attendance{}).Where("id = ?", aid).Take(&attendance).Error
	if err != nil {
		return &Attendance{}, err
	}
	if attendance.ID != 0 {
		err = db.Debug().Model(&Teacher{}).Where("id = ?", attendance.TeacherID).Take(&attendance.Teacher).Error
		if err != nil {
			return &Attendance{}, err
		}
		err = db.Debug().Model(&Student{}).Where("roll_no = ?", attendance.StudentRollNo).Take(&attendance.Student).Error
		if err != nil {
			return &Attendance{}, err
		}
		err = db.Debug().Model(&Faculty{}).Where("id = ?", attendance.Teacher.FacultyID).Take(&attendance.Teacher.Faculty).Error
		if err != nil {
			return &Attendance{}, err
		}
	}
	return attendance, nil
}

func (attendance *Attendance) SetAbsent(db *gorm.DB) (*Attendance, error) {
	var err error
	err = db.Debug().Model(&Attendance{}).Where("id = ?", attendance.ID).UpdateColumns(
		Attendance{
			Status: false,
		}).Error
	if err != nil {
		return &Attendance{}, err
	}
	if attendance.ID != 0 {
		err = db.Debug().Model(&Teacher{}).Where("id = ?", attendance.TeacherID).Take(&attendance.Teacher).Error
		if err != nil {
			return &Attendance{}, err
		}
		err = db.Debug().Model(&Student{}).Where("roll_no = ?", attendance.StudentRollNo).Take(&attendance.Student).Error
		if err != nil {
			return &Attendance{}, err
		}
		err = db.Debug().Model(&Faculty{}).Where("id = ?", attendance.Teacher.FacultyID).Take(&attendance.Teacher.Faculty).Error
		if err != nil {
			return &Attendance{}, err
		}
	}
	return attendance, nil
}

func (attendance *Attendance) SetPresent(db *gorm.DB) (*Attendance, error) {
	var err error
	err = db.Debug().Model(&Attendance{}).Where("id = ?", attendance.ID).UpdateColumns(
		Attendance{
			Status: true,
		}).Error
	if err != nil {
		return &Attendance{}, err
	}
	if attendance.ID != 0 {
		err = db.Debug().Model(&Teacher{}).Where("id = ?", attendance.TeacherID).Take(&attendance.Teacher).Error
		if err != nil {
			return &Attendance{}, err
		}
		err = db.Debug().Debug().Model(&Student{}).Where("roll_no = ?", attendance.StudentRollNo).Take(&attendance.Student).Error
		if err != nil {
			return &Attendance{}, err
		}
		err = db.Debug().Model(&Faculty{}).Where("id = ?", attendance.Teacher.FacultyID).Take(&attendance.Teacher.Faculty).Error
		if err != nil {
			return &Attendance{}, err
		}
	}
	return attendance, nil
}

func (attendance *Attendance) DeleteAttendance(db *gorm.DB, aid uint32, tid uint32) (*Attendance, error) {
	db = db.Debug().Model(&Attendance{}).Where("id = ? and teacher_id = ?", aid, tid).Take(&Attendance{}).Delete(&Attendance{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return &Attendance{}, db.Error
		}
		return &Attendance{}, db.Error
	}
	return attendance, nil
}
