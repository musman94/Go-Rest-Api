package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/musman94/Otsimo-Task/controller"
)

// Defines HTTP request routes
func main() {
	r := mux.NewRouter()

	//Candidate endpoints
	r.HandleFunc("/candidate/createCandidate", controller.CreateCandidate).Methods("POST")
	r.HandleFunc("/candidate/readCandidate/{_id}", controller.ReadCandidate).Methods("GET")
	r.HandleFunc("/candidate/deleteCandidate/{_id}", controller.DeleteCandidate).Methods("DELETE")
	r.HandleFunc("/candidate/arrangeMeeting/{_id}", controller.ArrangeMeeting).Methods("PATCH")
	r.HandleFunc("/candidate/completeMeeting/{_id}", controller.CompleteMeeting).Methods("PATCH")
	r.HandleFunc("/candidate/denyCandidate/{_id}", controller.DenyCandidate).Methods("PATCH")
	r.HandleFunc("/candidate/acceptCandidate/{_id}", controller.AcceptCandidate).Methods("PATCH")

	//Assignee endpoints
	r.HandleFunc("/assignee/findAssigneeIDByName/{name}", controller.FindAssigneeIDByName).Methods("GET")
	r.HandleFunc("/assignee/findAssigneesCandidates/{_id}", controller.FindAssigneesCandidates).Methods("GET")
	//r.HandleFunc("/createAssignee", controller.CreateAssignee).Methods("POST")

	//Starts the server
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
