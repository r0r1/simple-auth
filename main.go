package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"bitbucket.org/rorikurniadi/rori_kurniadi_test/models"
	"bitbucket.org/rorikurniadi/rori_kurniadi_test/resources"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Hello Handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Simple Auth API!\n"))
}

// JWT Header
func jwtHeader(req *http.Request, key string) (string, error) {
	authHeader := req.Header.Get(key)

	if authHeader == "" {
		return "", errors.New("auth header empty")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("invalid auth header")
	}

	return parts[1], nil
}

// middleware to protect private pages
func Middleware(page http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var tokenString string
		var err error
		tokenString, err = jwtHeader(req, "Authorization")

		if err != nil {
			json.NewEncoder(res).Encode(err.Error())
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &resources.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("JWT_TOKEN"), nil
		})

		if claims, ok := token.Claims.(*resources.MyCustomClaims); ok && token.Valid {
			res.Header().Set("CURRENT_USER_EMAIL", claims.Email)
			res.Header().Set("CURRENT_USER_ID", claims.Id)
			page.ServeHTTP(res, req)
		} else {
			res.Header().Set("Content-Type", "application/json; charset=UTF-8")
			res.WriteHeader(401)
			json.NewEncoder(res).Encode(&resources.JsonError{Message: "Unauthorized"})
			return
		}
	})
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	resource := resources.ResourceDB(db)

	// router
	r := mux.NewRouter()
	r.HandleFunc("/", HelloHandler).Methods("GET")
	r.HandleFunc("/api/register", resource.Register).Methods("POST")
	r.HandleFunc("/api/auth", resource.Auth).Methods("POST")
	r.HandleFunc("/api/forgot-password", resource.ForgotPassword).Methods("POST")

	// linkedin OAuth
	r.HandleFunc("/api/linkedin_url", resource.GetLinkedinURL).Methods("GET")
	r.HandleFunc("/api/linkedin_callback", resource.LinkedinCallback).Methods("GET")

	// jwtMiddleware
	r.HandleFunc("/api/authenticated", Middleware(resource.Authenticated)).Methods("GET")
	r.HandleFunc("/api/users/{user_id}", Middleware(resource.UpdateUser)).Methods("PATCH")
	// Bind to a port and pass our router in

	headersOk := handlers.AllowedHeaders([]string{"Origin", "Authorization", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	exposeHeaders := handlers.ExposedHeaders([]string{""})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH"})
	credentials := handlers.AllowCredentials()
	log.Fatal(http.ListenAndServe(":8888", handlers.CORS(originsOk, headersOk, exposeHeaders, credentials, methodsOk)(r)))
}
