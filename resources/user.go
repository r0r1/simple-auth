package resources

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/rorikurniadi/rori_kurniadi_test/models"

	"github.com/gorilla/mux"
)

// UpdateUser
func (r *Resource) UpdateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(req.Body)
	var user models.User
	err := decoder.Decode(&user)

	if err != nil {
		log.Print(err.Error())
	}

	vars := mux.Vars(req)
	userID := vars["user_id"]
	postParam := user

	if r.db.First(&user, userID).RecordNotFound() {
		res.WriteHeader(404)
		json.NewEncoder(res).Encode(JsonError{Message: "User Not Found."})
		return
	}

	user.Name = postParam.Name
	user.Title = postParam.Title
	user.Contact = postParam.Contact
	user.Address = postParam.Address
	user.Latitude = postParam.Latitude
	user.Longitude = postParam.Longitude
	r.db.Save(&user)

	json.NewEncoder(res).Encode(user)
}
