package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/gorilla/mux"

	config "api/config"
)

func UpdateCake(w http.ResponseWriter, r *http.Request) {

	var response Response
	var responses ResponseCon
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

		sql := `UPDATE cake SET `
		sqlParts := make([]string, 0, 4)
		args := make([]interface{}, 0, 4)

		if cake.Title != "" {
			sqlParts = append(sqlParts, `title = ?`)
			args = append(args, cake.Title)
		}

		if cake.Description != "" {
			sqlParts = append(sqlParts, `description = ?`)
			args = append(args, cake.Description)
		}
		//temporarry nil = 0
		if cake.Rating != 0 {
			sqlParts = append(sqlParts, `rating = ?`)
			args = append(args, cake.Rating)
		}
		if cake.Image != "" {
			sqlParts = append(sqlParts, `image = ?`)
			args = append(args, cake.Image)
		}

		sql += strings.Join(sqlParts, `,`) + ` WHERE id = ?`
		args = append(args, id)

		//inserting to database
		db := config.Connect()
		defer db.Close()
		_, err := db.Exec(sql, args...)

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

}
