package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	config "api/config"
)

func ListOfCake(w http.ResponseWriter, r *http.Request) {
	var cake Cake
	var response Response
	var listOfCake []Cake

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id,title,description,rating,image,created_at,updated_at FROM cake ORDER BY rating DESC, title ASC")

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

	response.ResponseCode = 200
	response.Total = len(listOfCake)
	response.Message = "Success"
	response.Data = listOfCake

	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
