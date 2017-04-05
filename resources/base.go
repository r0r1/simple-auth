package resources

import (
	"strconv"
	"time"

	"bitbucket.org/rorikurniadi/rori_kurniadi_test/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
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
	claims := &MyCustomClaims{
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Id:        strconv.Itoa(int(user.ID)),
			Issuer:    "localhost:8000",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", false
	}
	return jwtToken, true
}
