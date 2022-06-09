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

func DetailOfCake(w http.ResponseWriter, r *http.Request) {
	var cake Cake
	var response Response
	var listOfCake []Cake
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
		//continue checking to database
		db := config.Connect()
		defer db.Close()

		rows, err := db.Query("SELECT * FROM cake WHERE id=?", id)

		if err != nil {
			log.Print(err)
		}

		for rows.Next() {
			err = rows.Scan(
				&cake.Id,
				&cake.Title,
				&cake.Description,
				&cake.Rating,
				&cake.Image,
				&cake.Created_at,
				&cake.Updated_at,
			)
			if err != nil {
				response.ResponseCode = 500
				response.Message = fmt.Sprintf("Internal Server Error %v", err)
				log.Fatal(err.Error())
			} else {
				listOfCake = append(listOfCake, cake)
			}
		}

		if len(listOfCake) == 0 {
			response.ResponseCode = 202
			response.Total = len(listOfCake)
			response.Message = fmt.Sprintf(`Cake with ID %v doesn't exist `, id)
			json.NewEncoder(w).Encode(response)
		} else {
			json.NewEncoder(w).Encode(cake)
		}

	}

}
