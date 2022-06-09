package api

import "net/http"

type (
	// Service interface
	Service interface {
		ListOfCake(w http.ResponseWriter, r *http.Request)
		DetailOfCake(w http.ResponseWriter, r *http.Request)
		AddNewCake(w http.ResponseWriter, r *http.Request)
		UpdateCake(w http.ResponseWriter, r *http.Request)
		DeleteCake(w http.ResponseWriter, r *http.Request)
	}

	// Cake struct
	Cake struct {
		Id          int     `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Rating      float32 `json:"rating"`
		Image       string  `json:"image"`
		Created_at  string  `json:"created_at"`
		Updated_at  string  `json:"updated_at"`
	}

	// Response struct
	Response struct {
		ResponseCode int    `json:"responsecode"`
		Total        int    `json:"total"`
		Data         []Cake `json:"data"`
		Message      string `json:"message"`
	}

	// ResponseCon struct
	ResponseCon struct {
		ResponseCode int    `json:"responsecode"`
		Message      string `json:"message"`
	}
)
