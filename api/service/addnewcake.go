package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	config "api/config"
)

func AddNewCake(w http.ResponseWriter, r *http.Request) {
	var response Response
	var responses ResponseCon
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	decoder := json.NewDecoder(r.Body)
	var cake Cake

	errorz := decoder.Decode(&cake)
	if errorz != nil {
		response.ResponseCode = 400
		response.Message = "Invalid request payload"
		json.NewEncoder(w).Encode(response)
		log.Print(errorz)
		return
	}

	//checking body for require variable (assumption: rating not require and other require)
	if cake.Title == "" {
		response.ResponseCode = 401
		response.Message = "No Title Specified"
		json.NewEncoder(w).Encode(response)
		return
	}
	if cake.Description == "" {
		response.ResponseCode = 402
		response.Message = "No Description Specified"
		json.NewEncoder(w).Encode(response)
		return
	}
	if cake.Image == "" {
		response.ResponseCode = 403
		response.Message = "No Image Specified"
		json.NewEncoder(w).Encode(response)
		return
	}

	//inserting to database
	db := config.Connect()
	defer db.Close()
	_, err := db.Exec("INSERT INTO cake (title, description, rating, image) VALUES (? , ? , ? , ?)", cake.Title, cake.Description, cake.Rating, cake.Image)

	if err != nil {
		log.Print(err)
		response.ResponseCode = 500
		response.Message = fmt.Sprintf("error %s", err)
		json.NewEncoder(w).Encode(response)
	} else {
		responses.ResponseCode = 200
		responses.Message = "Success"
		json.NewEncoder(w).Encode(responses)
	}
}
