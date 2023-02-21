package helper

import (
	"errors"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), 10)
}

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be blank")
	}
	if len(password) < 8 || len(password) > 30 {
		return errors.New("password length should be 8 to 30 characters")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("^[A-Za-z0-9$_@.#]+$"))) != nil {
		return errors.New("password should contain only alphabetic characters, numbers and special characters(@, $, _, ., #)")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("[0-9]"))) != nil {
		return errors.New("password should contain at least one number")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("[A-Za-z]"))) != nil {
		return errors.New("password should contain at least one alphabetic character")
	}
	return nil
}

func ValidateLogin(login string) error {
	if login == "" {
		return errors.New("login cannot be blank")
	}
	if len(login) < 5 || len(login) > 15 {
		return errors.New("login length should be 5 to 15 characters")
	}
	if validation.Validate(login, validation.Match(regexp.MustCompile("^[A-Za-z0-9$@_.#]+$"))) != nil {
		return errors.New("login should contain only alphabetic characters, numbers and special characters(@, $, _, ., #)")
	}
	return nil
}

func ValidateDate(date string) error {
	if date == "" {
		return errors.New("date is blank")
	}

	if validation.Validate(date, validation.Date("02-01-2006")) != nil {
		return errors.New("date must be DD-MM-YYYY format")
	}
	return nil
}

func ValidatePhoneNumber(phoneNumber string) error {
	if phoneNumber == "" {
		return errors.New("phone_number is blank")
	}
	pattern := regexp.MustCompile(`^(\+[0-9]{12})$`)

	if !(pattern.MatchString(phoneNumber) && phoneNumber[0:4] == "+998") {
		return errors.New("phone_number must be +998XXXXXXXXX")
	}
	return nil
}

func ValidateIp(ip string) error {
	if validation.Validate(ip, is.IPv4) != nil {
		return errors.New("ip must be in IPv4 Form")
	}

	return nil
}

func ValidatePort(port string) error {
	return validation.Validate(port, validation.Required, validation.Length(1, 5), is.Digit)
}

func ValidateOrderNo(orderNo int32) error {
	if orderNo < 0 {
		return errors.New("Order Number should be positive")
	}

	return nil
}
func DaysBetween(a, b time.Time) int64 {
	days := (b.Unix() - a.Unix()) / (60 * 60 * 24)
	return days
}

func CheckingDateRange(beginDate, endDate string) error {
	t1, _ := time.Parse("2006-01-02", beginDate)
	t2, _ := time.Parse("2006-01-02", endDate)
	days := DaysBetween(t1, t2)

	if days > 31 {
		return errors.New("time range is invalid. You can only get monthly reports")
	}

	return nil

}
func DateFormatting(date string) string {
	d, _ := time.Parse("02-01-2006", date)
	formattedDate := d.Format("2006-01-02")
	return formattedDate

}

func ValidPinfl(pinfl string) error {
	if pinfl == "" {
		return errors.New("error client passport_pinfl requirement body to model")
	}
	pattern := regexp.MustCompile(`^([0-9]{14})$`)

	if !(pattern.MatchString(pinfl)) {
		return errors.New("phone_number must be 14 digits")
	}
	return nil
}
