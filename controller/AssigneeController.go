package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/musman94/Otsimo-Task/service"

	. "github.com/musman94/Otsimo-Task/helper"
	. "github.com/musman94/Otsimo-Task/model"
)

//FindAssigneeIDByName Endpoint
//Fetches the id of an assignee using the name by calling the service layer function
//Returns the id
func FindAssigneeIDByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	result := service.FindAssigneeIDByName(name)

	if result == "" {
		RespondWithError(w, http.StatusNotFound, "Assignee with the given name doesnot exist")
		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"Assignee's ID": result})
}

//FindAssigneesCandidates Endpoint
//Gets all of the candidates of an assignee by calling the service layer function
//Returns the candidates if succesful, an error otherwise
func FindAssigneesCandidates(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_id := params["_id"]

	result, err := service.FindAssigneesCandidates(_id)

	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Invalid Assignee ID")
		return
	}

	RespondWithJson(w, http.StatusOK, map[string][]Candidate{"Candidates": result})
}   

