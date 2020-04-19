package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/musman94/Otsimo-Task/service"

	. "github.com/musman94/Otsimo-Task/helper"
	. "github.com/musman94/Otsimo-Task/model"
)

//CreateCandidate Endpoint
//Creates a candidate by calling the service layer function
//Returns the candidate if succesful, an error otherwise
func CreateCandidate(w http.ResponseWriter, r *http.Request) {
	var candidate Candidate

	if json.NewDecoder(r.Body).Decode(&candidate) != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result, err := service.CreateCandidate(candidate)

	if err != nil {
		errVal := err.(*HttpError)
		RespondWithError(w, errVal.Code, errVal.Msg)
		return
	}

	RespondWithJson(w, http.StatusCreated, result)
}

//ReadCandidate Endpoint
//Gets the candidate by calling the service layer function
//Returns the candidate if succesful, an error otherwise
func ReadCandidate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var _id = params["_id"]

	result, err := service.ReadCandidate(_id)

	if err != nil {
		errVal := err.(*HttpError)
		RespondWithError(w, errVal.Code, errVal.Msg)
		return
	}

	RespondWithJson(w, http.StatusOK, result)
}

//DeleteCandidate Endpoint
//Deletes the candidate by calling the service layer function
//Returns the id of the candidate if succesful, an error otherwise
func DeleteCandidate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_id := params["_id"]

	err := service.DeleteCandidate(_id)

	if err != nil {
		errVal := err.(*HttpError)
		RespondWithError(w, errVal.Code, errVal.Msg)
		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"Candidate deleted": _id})
}

//ArrangeMeeting Endpoint
//Arranges a meeting with the candidate by calling the service layer function
//Returns the id of the candidate if succesful, an error otherwise
func ArrangeMeeting(w http.ResponseWriter, r *http.Request) {
	var candidate Candidate

	params := mux.Vars(r)
	_id := params["_id"]

	if json.NewDecoder(r.Body).Decode(&candidate) != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := service.ArrangeMeeting(_id, candidate.NextMeeting)

	if err != nil {
		errVal := err.(*HttpError)
		RespondWithError(w, errVal.Code, errVal.Msg)
		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"Meeting arranged with candidate": _id})
}

//CompleteMeeting Endpoint
//Completes a meeting with the candidate by calling the service layer function
//Returns the id of the candidate if succesful, an error otherwise
func CompleteMeeting(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_id := params["_id"]

	err := service.CompleteMeeting(_id)

	if err != nil {
		errVal := err.(*HttpError)
		RespondWithError(w, errVal.Code, errVal.Msg)
		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"Meeting completed with candidate": _id})
}

//DenyCandidate Endpoint
//Denies a candidate by calling the service layer function
//Returns the id of the candidate if succesful, an error otherwise
func DenyCandidate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_id := params["_id"]

	err := service.DenyCandidate(_id)

	if err != nil {
		errVal := err.(*HttpError)
		RespondWithError(w, errVal.Code, errVal.Msg)
		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"Candidate denied": _id})
}

//AcceptCandidate Endpoint
//Accepts a candidate by calling the service layer function
//Returns the id of the candidate if succesful, an error otherwise
func AcceptCandidate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_id := params["_id"]

	err := service.AcceptCandidate(_id)

	if err != nil {
		errVal := err.(*HttpError)
		RespondWithError(w, errVal.Code, errVal.Msg)
		return
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"Candidate accepted": _id})
}
