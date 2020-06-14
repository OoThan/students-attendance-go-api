package formaterrors

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "acronym") {
		return errors.New("Acronym Already Taken ")
	}
	if strings.Contains(err, "faculty_name") {
		return errors.New("Faculty Name Already Taken ")
	}
	if strings.Contains(err, "faculty_mail") {
		return errors.New("Faculty Mail Already Taken ")
	}
	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken ")
	}
	if strings.Contains(err, "roll_no") {
		return errors.New("RollNo already taken ")
	}
	if strings.Contains(err, "status") {
		return errors.New("Required Status ")
	}
	return errors.New("Incorrect Details ")
}
