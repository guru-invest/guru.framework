package validations

import (
	"regexp"
	"strings"
)

func ValidateEmail(email string) bool {
	rEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	res := rEmail.MatchString(email)
	return res
}

func ValidateDocumentNumber(document_number string) bool {
	rDocument := regexp.MustCompile("^[0-9]*$")
	res := rDocument.MatchString(document_number)
	if res && strings.Count(document_number, "") >= 12 {
		return res
	} else {
		return false
	}
}

func ValidateContactNumber(contact string) bool {
	rContact := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	res := rContact.MatchString(contact)
	if res && strings.Count(contact, "") <= 12 {
		return res
	} else {
		return false
	}
}