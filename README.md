# Otsimo-Task
## Running the server
To run the server type `go run server.go` in the terminal. The server will start running on http://127.0.0.1:3000. 

## Endpoints
You can read more about each of the endpoints by clicking on their names. 

|Name                                                                    |route                                    |method |                                                  
|------------------------------------------------------------------------|-----------------------------------------|-------                  
[createCandidate](examples/candidate/createCandidate.md)                 |/candidate/createCandidate               |POST   |
[readCandidate](examples/candidate/readCandidate.md)                     |/candidate/readCandidate/{_id}           |GET    |
[deleteCandidate](examples/candidate/deleteCandidate.md)                 |/candidate/deleteCandidate/{_id}         |DELETE |
[arrangeMeeting](examples/candidate/arrangeMeeting.md)                   |/candidate/arrangeMeeting/{_id}          |PUT    |
[completeMeeting](examples/candidate/completeMeeting.md)                 |/candidate/completeMeeting/{_id}         |PUT    |
[denyCandidate](examples/candidate/denyCandidate.md)                     |/candidate/denyCandidate/{_id}           |PUT    |
[acceptCandidate](examples/candidate/acceptCandidate.md)                 |/candidate/acceptCandidate/{_id}         |PUT    |
[findAssigneesCandidates](examples/candidate/findAssigneesCandidates.md) |/candidate/findAssigneesCandidates/{_id} |GET    |
[findAssigneeIDByName](examples/assignee/findAssigneeIDByName.md)        |/assignee/findAssigneeIDByName{name}     |GET    |

## Testing
For testing you can use a tool like POSTMAN and call each of the endpoints with the required parameters.  

## Structure
Layers:        
Controller   
Service   
Dao   

The Controller layer is mainly responsible for getting the results and returning them in JSON format. The Service layer checks the inputs/parameters, makes changes to the objects and forwards them to the Dao layer. The Dao layer acts as a gateway to the database.   

`server.go`is the entry point of the server. It calls either `CandidateController.go` or `AssigneeController.go` depending on the route. They in turn call either `CandidateService.go` or `AssigneeService.go`. which forward the request to either `CandidateDAO.go` or `AssigneeDAO.go`.   

Layer specific files can be found in the folder with the layer's name e.g Controller files are in a folder named controller.   
helper dir contains a file having some helper functions that are used to connect to the DB, get JSON responses etc.   
model dir contains the model definitions for both the Candidate and Assignee.   
dump dir contains the database `otsimo_db` that I used while working on the project.   
examples dir contains the api documentation.   

## Get the external modules

 ```sh
go get go.mongodb.org/mongo-driver gopkg.in/mgo.v2
```

## Clone the repo and run

$ git clone https://github.com/musman94/Otsimo-Task.git
$ go run server.go
