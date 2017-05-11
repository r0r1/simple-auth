package resources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rorikurniadi/simple-auth/models"

	config "github.com/rorikurniadi/simple-auth/configs"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/linkedin"
)

var dbConfig = config.ReadConfig()

//state should be regenerated per auth request
var (
	State = "linkedin_rori_kurniadi_simple_auth"
)

// GET LINKEDIN URL ACCESS
func (r *Resource) GetLinkedinURL(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	inurl := GetConfig().AuthCodeURL(State)

	json.NewEncoder(res).Encode(JsonSuccess{Message: inurl})
}

// Callback Linkedin
func (r *Resource) LinkedinCallback(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//check state validity, see url := Config.AuthCodeURL(state) above
	stateCheck := req.FormValue("state")
	if State != stateCheck {
		http.Error(res, fmt.Sprintf("Wrong state string: Expected %s, got %s. Please, try again", State, stateCheck), http.StatusBadRequest)
		return
	}

	token, err := GetConfig().Exchange(oauth2.NoContext, req.FormValue("code"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	client := GetConfig().Client(oauth2.NoContext, token)
	request, err := http.NewRequest("GET", "https://api.linkedin.com/v1/people/~:(email-address,first-name,last-name,id,headline)?format=json", nil)
	if err != nil {
		println("egt eeee")
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	request.Header.Set("Bearer", token.AccessToken)
	response, err := client.Do(request)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	defer response.Body.Close()
	str, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.User
	var inuser struct {
		Id        string
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Headline  string
		Email     string `json:"emailAddress"`
	}

	err = json.Unmarshal(str, &inuser)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	data := r.db.Where("email = ?", inuser.Email).Find(&user)
	if data.RecordNotFound() {
		user.Name = inuser.FirstName + " " + inuser.LastName
		user.Title = inuser.Headline
		user.Email = inuser.Email

		if err := r.db.Save(&user).Error; err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		result, valid := CreateClaim(&user)
		if !valid {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}

		http.Redirect(res, req, dbConfig.APP_URL+"oauth?token="+result+"&type=signup", 301)
		return
	}

	result, valid := CreateClaim(&user)
	if !valid {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(res, req, dbConfig.APP_URL+"oauth?token="+result+"&type=login", 301)
}

//config
func GetConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     dbConfig.APP_ID,           // change this to yours
		ClientSecret: dbConfig.APP_SECRET,       //change this to yours
		RedirectURL:  dbConfig.APP_CALLBACK_URL, // change this to your webserver adddress
		Scopes:       []string{"r_basicprofile", "r_emailaddress"},
		Endpoint:     linkedin.Endpoint,
	}
}
