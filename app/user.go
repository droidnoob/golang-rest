package app

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Claim object for the JWT Token
type Token struct {
	Email string
	jwt.StandardClaims
}

// User model
type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"-" form:"token" sql:"-"`
}

func (user *User) Validate() (string, bool) {
	if !strings.Contains(user.Email, "@") {
		return "Email is invalid", false
	}
	if len(user.Password) < 8 {
		return "Password needs to be atleast 8 characters", false
	}

	var temp User

	err := GetDB().Table("users").Where("email=?", user.Email).First(&temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "Error connecting to the database", false
	}
	if temp.Email != "" {
		return "Email already exists!", false
	}
	return "Validated", true

}

func (user *User) CreateToken() {
	userToken := &Token{Email: user.Email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), userToken)
	tokenString, _ := token.SignedString([]byte(os.Getenv("jwt_secret_password")))
	user.Token = tokenString
}

func (user *User) hashPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
}

func (user *User) Create() (string, bool) {

	msg, ok := user.Validate()
	if !ok {
		return msg, false
	}
	user.hashPassword()
	GetDB().Create(user)

	if user.ID <= 0 {
		return "Failed to create account", false
	}

	user.CreateToken()
	return user.Token, true
}

func Login(email string, password string) (string, bool) {
	user := &User{}
	err := GetDB().Table("users").Where("email=?", email).First(user).Error

	if err != nil {
		return "Invalid Login or Password", false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "Invalid Login or Password", false
	}

	user.Password = ""

	user.CreateToken()
	return user.Token, true
}

func GetUser(email interface{}) (*User, bool) {
	user := &User{}

	GetDB().Table("users").Where("email=?", email).First(user)
	if user.Email == "" {
		return user, false
	}

	user.Password = ""
	return user, true

}

func (user *User) EditUser(email interface{}) (string, bool) {
	if user.Password != "" && !(len(user.Password) < 8) {
		user.hashPassword()
	}
	err := GetDB().Model(&user).Where("email=?", email).Updates(User{Name: user.Name, Password: user.Password}).Error
	if err != nil {
		return "Error during Update.", false
	}
	user.Password = ""

	return "Succesfully updated", true

}
