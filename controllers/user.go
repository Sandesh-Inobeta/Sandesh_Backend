package controllers

import (
	"encoding/json"
	"net/http"
	helpers "sandesh/helper"
	"sandesh/interfaces"
)

func GetName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var nameReq interfaces.NameRequest
	err := json.NewDecoder(r.Body).Decode(&nameReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, map[string]string{
		"Name": nameReq.Name,
	})
}
