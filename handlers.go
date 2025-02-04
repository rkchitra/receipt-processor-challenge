/*
this file contains the handlers for both HTTP requests and calls appropriate
functions to service them; the handlers are responsible for interpreting the incoming request and constructing the JSON response
*/

package main

import (
"net/http"
"github.com/gorilla/mux"
"encoding/json"
"strings"
)

/*
this function handles the POST request when a receipt is submitted. It examines
the receipt, performs validation, and either allocates points and assigns an Id or returns the appropriate error message
*/
func processReceiptsHandler(w http.ResponseWriter, r *http.Request) {

var receipt Receipt 

//extract receipt details, if it's not in the predefined format, return error
if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
WriteErrorResponse(w, http.StatusBadRequest, errMsg) 
return
}

// validation 

//validate if retailer's name is present
if len(receipt.Retailer) < 1 || len(receipt.Items) < 1{
WriteErrorResponse(w, http.StatusBadRequest, errMsg)
return
}

//call function to allocate points for the receipt; it returns errors if
//there are any issues with the payload data
id, err, statusCode := ProcessReceipt(receipt)

//error message returned if any errors present
if err != nil {
WriteErrorResponse(w, statusCode, err.Error())
return
}

//store id in response
id = strings.TrimSpace(id)

//return id in JSON format
response := ReceiptResponse{Id: id}

//embed details in JSON response
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(response)
}

/*
function to handle the receipt id received and either return points if valid
or appropriate error message if invalid
*/
func retrieveReceiptsHandler(w http.ResponseWriter, r *http.Request) {

//extract receipt ID from URL
vars := mux.Vars(r)
id := vars["id"]

//function call to get number of points allotted to receipt id if valid
points, err, statusCode := RetrieveReceipts(id)

//if called function returns errors, write the same in the response and return
if err != nil {
WriteErrorResponse(w, statusCode, err.Error())
return
}

//write the points allotted in JSON format and return response
response := PointsResponse {
Points: points,
}

//embed points details in the JSON response
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(response)
}

//function to write error response
func WriteErrorResponse(w http.ResponseWriter, statusCode int, err string) {
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(statusCode)

//create response of type ErrorResponse with message embedded within it
response := ErrorResponse{Error:err}
json.NewEncoder(w).Encode(response)
}
