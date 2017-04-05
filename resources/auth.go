package resources

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/rorikurniadi/rori_kurniadi_test/middleware"
	"bitbucket.org/rorikurniadi/rori_kurniadi_test/models"

	"golang.org/x/crypto/bcrypt"

	"fmt"
)

// Register
func (ar *Resource) Register(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(req.Body)
	var user models.User
	err := decoder.Decode(&user)

	if err != nil {
		log.Print(err.Error())
	}
	// validation
	if message, err := middleware.AuthValidate(user, res, req); err == false {
		res.WriteHeader(422)
		json.NewEncoder(res).Encode(JsonError{Message: message})
		return
	}

	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashedPassword)
	if err := ar.db.Save(&user).Error; err != nil {
		json.NewEncoder(res).Encode(JsonError{Message: "Register failed."})
		return
	}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(res).Encode(user); err != nil {
		panic(err)
	}
}

// Auth
func (ar *Resource) Auth(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(req.Body)
	var user models.User
	err := decoder.Decode(&user)

	if err != nil {
		log.Print(err.Error())
	}

	// validation
	if message, err := middleware.AuthValidate(user, res, req); err == false {
		res.WriteHeader(422)
		json.NewEncoder(res).Encode(JsonError{Message: message})
		return
	}

	inputPassword := []byte(user.Password)
	data := ar.db.Where("email = ?", user.Email).Find(&user)

	if data.RecordNotFound() {
		res.WriteHeader(400)
		json.NewEncoder(res).Encode(JsonError{Message: "Email is invalid."})
		return
	}

	hashedPassword := []byte(user.Password)
	checkPassword := bcrypt.CompareHashAndPassword(hashedPassword, inputPassword)

	if checkPassword != nil {
		res.WriteHeader(400)
		json.NewEncoder(res).Encode(JsonError{Message: "Password is invalid."})
		return
	}

	// CreateClaim
	token, valid := CreateClaim(&user)
	if !valid {
		json.NewEncoder(res).Encode(JsonError{Message: "Failed generate token"})
	}
	json.NewEncoder(res).Encode(JsonSuccess{Message: token})
}

// Authenticated
func (r *Resource) Authenticated(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var user models.User
	data := r.db.First(&user, res.Header().Get("CURRENT_USER_ID"))

	if data.RecordNotFound() {
		res.WriteHeader(404)
		json.NewEncoder(res).Encode(JsonError{Message: "User not found."})
		return
	}

	json.NewEncoder(res).Encode(user)
}

// Forgot Password function
func (r *Resource) ForgotPassword(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(req.Body)
	var user models.User
	err := decoder.Decode(&user)

	if err != nil {
		log.Print(err.Error())
	}

	if message, err := middleware.AuthValidate(user, res, req); err == false {
		res.WriteHeader(422)
		json.NewEncoder(res).Encode(JsonError{Message: message})
		return
	}

	password := []byte(user.Password)
	data := r.db.Where("email = ?", user.Email).Find(&user)

	if data.RecordNotFound() {
		res.WriteHeader(404)
		json.NewEncoder(res).Encode(JsonError{Message: "Email not found."})
		return
	}

	fmt.Printf("Password: %s", password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashedPassword)
	r.db.Save(&user)
	json.NewEncoder(res).Encode(JsonSuccess{Message: "Forgot password successfull."})
}
