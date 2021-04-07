package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
	"unicode"
)

//swagger:model user
type User struct {
	Id         int64     `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	SurName    string    `json:"surName" gorm:"sur_name"`
	Password   string    `json:"password"`
	Login      string    `json:"login"`
	Email      string    `json:"email"`
	ClaimToken string    `json:"claimToken" gorm:"-"`
}

func (User) TableName() string {
	return fmt.Sprintf("%s.%s", SchemaBooking, "user")
}

type UserClaims struct {
	User
	jwt.StandardClaims
}

func (u *User) Validate() error {
	if len(u.Login) == 0 {
		return fmt.Errorf("обязательное поле: логин")
	}
	if !strings.Contains(u.Email, "@") {
		return fmt.Errorf("почта должна содержать @")
	}

	if !strings.Contains(strings.Split(u.Email, "@")[1], ".") {
		return fmt.Errorf("почта содержать домен после @ (mail.ru)")
	}
	if len(u.Password) == 0 {
		return fmt.Errorf("обязательное поле: пароль")
	}
	err := u.VerifyPassword(u.Password)
	if err != nil {
		return err
	}

	temp := &User{}

	userTX := GetDB().
		Where("login = ?", u.Login).
		First(temp)

	if userTX.Error != nil && userTX.Error != gorm.ErrRecordNotFound {
		return fmt.Errorf("%s: %v, ", ErrorUnexpected, userTX.Error)
	}
	if temp.Login != "" || userTX.RowsAffected > 0 {
		return fmt.Errorf("пользователь с таким логином уже существует")
	}
	return nil
}

func (u *User) LoginCheck(cred BodyCredentials) error {
	if len(cred.Login) > 12 || len(cred.Login) < 11 || len(cred.Password) == 0{
		return fmt.Errorf("%s", ErrorSigIn)
	}
	if len(cred.Login) == 12 {
		if cred.Login[:1] == "+" {
			cred.Login = cred.Login[1:]
		}
	}
	if len(cred.Login) == 11 {
		if cred.Login[:1] == "8" {
			cred.Login =  "7" + cred.Login[1:]
		}
	}
	err := GetDB().
		Where("login = ?", cred.Login).
		First(&u).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("%s: %v, ", ErrorUnexpected, err.Error())
	}
	if err == gorm.ErrRecordNotFound {
		return fmt.Errorf("%s", ErrorSigIn)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(cred.Password))
	if err != nil {
		return fmt.Errorf("%s", ErrorSigIn)
	}
	u.Password = ""
	return nil
}

func (u *User) VerifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int
	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	if !lowercasePresent {
		return fmt.Errorf("пароль должен содержать хотя бы 1 символ в нижнем регистре")
	}
	if !uppercasePresent {
		return fmt.Errorf("пароль должен содержать хотя бы 1 символ в верхнем регистре")
	}
	if !numberPresent {
		return fmt.Errorf("пароль должен содержать хотя бы 1 номер")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		return fmt.Errorf(fmt.Sprintf("длинна пароля должна быть между %d и %d символами", minPassLength, maxPassLength))
	}
	return nil
}

func (u User) CreateJwtToken() (tokenString string, err error) {
	claims := UserClaims{
		u,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(12) * time.Hour).Unix(), //12 часов
			Issuer:    "booking",
			Subject:   fmt.Sprintf("%s", u.Login),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(os.Getenv("token_secret")))
	if err != nil {
		return
	}
	return
}

