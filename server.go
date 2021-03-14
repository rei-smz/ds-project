package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("content-type","application/json")
	
	query := r.URL.Query()
	switch r.URL.Path {
		case "/api/check-service":
			w.Write([]byte("ok"))
		case "/api/get-location":

		case "/api/get-facility-nearby":

		case "/api/add-navigation":

		case "/api/start-simulation":

		case "/api/pause-simulation":

	}
}
