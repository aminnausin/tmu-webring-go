package account

import (
	"log"
	"net/http"
	"tmu-webring-go/api"

	"github.com/gorilla/schema"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {
	var params = api.AccountParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Default().Println("ERROR: ", err)
		api.InternalErrorHandler(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")

}
func UpdateAccount(w http.ResponseWriter, r *http.Request) {

}
func CreateAccount(w http.ResponseWriter, r *http.Request) {

}
func DeleteAccount(w http.ResponseWriter, r *http.Request) {

}
