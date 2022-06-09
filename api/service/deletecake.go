package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"unicode"

	"github.com/gorilla/mux"

	config "api/config"
)

func DeleteCake(w http.ResponseWriter, r *http.Request) {
	var response ResponseCon
	var count int
	check := true
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	pathid := mux.Vars(r)
	id := pathid["id"]

	// checking id only contain number (because id stored as int in database)
	for _, c := range id {
		if !unicode.IsDigit(c) {
			check = false
		}
	}

	if !check {
		response.ResponseCode = 400
		response.Message = "id must only contain number"
		json.NewEncoder(w).Encode(response)
	} else {

		db := config.Connect()
		defer db.Close()

		rows, _ := db.Query("SELECT COUNT(*) FROM cake WHERE id=?", id)
		rows.Scan(&count)
		if count == 0 {
			response.ResponseCode = 202
			response.Message = fmt.Sprintf(`Cake with ID %v doesn't exist `, id)
		} else {
			_, err := db.Query("DELETE FROM cake WHERE id = ?", id)

			if err != nil {
				log.Print(err)
				response.ResponseCode = 500
				response.Message = "Internal Server Error"
			} else {
				response.ResponseCode = 200
				response.Message = "Success"
			}

		}

		json.NewEncoder(w).Encode(response)
	}

}
