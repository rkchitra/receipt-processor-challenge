//this file is the entry point and defines endpoints and corresponding handlers

package main

import (
"fmt"
"log"
"net/http"
"github.com/gorilla/mux"
)

//entry point function that defines the endpoints
func main() {
//set up endpoints and corresponding function handlers

r := mux.NewRouter()

//handlers for both endpoints: POST and GET methods 

r.HandleFunc("/receipts/process", processReceiptsHandler).Methods("POST")
r.HandleFunc("/receipts/{id}/points", retrieveReceiptsHandler).Methods("GET")

//announce on console that server is up and running and log messages from 8080
fmt.Println("\nServer running on port 8080")
log.Fatal(http.ListenAndServe(":8080", r))
}
