package resources

import (
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	config "github.com/rorikurniadi/simple-auth/configs"
	"github.com/rorikurniadi/simple-auth/models"
)

// Auth DB initializes the storage
func ResourceDB(db *gorm.DB) *Resource {
	return &Resource{db}
}

var mySigningKey = []byte("JWT_TOKEN")

// Resource
type Resource struct {
	db *gorm.DB
}

type JsonError struct {
	Message string `json:"error"`
}

type JsonSuccess struct {
	Message string `json:"data"`
}

type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Create the Claims
func CreateClaim(user *models.User) (string, bool) {
	var config = config.ReadConfig()
	claims := &MyCustomClaims{
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Id:        strconv.Itoa(int(user.ID)),
			Issuer:    config.APP_URL,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", false
	}
	return jwtToken, true
}
