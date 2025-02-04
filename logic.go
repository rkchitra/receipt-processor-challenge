/*
this file contains the necessary functions to handle the underlying logic for
each endpoint
*/  

package main

import(
"fmt"
"strconv"
"math"
"strings"
"unicode"
"os/exec"
) 

//global map used by the methods in this file to map the receipt id to points 
var ReceiptsMapper = make(map[string]int)

//constants for status codes
const badRequest, internalServerError, notFound, success = 400, 500, 404, 200

//constant that stores error message for data issues
const errMsg = "The receipt is invalid."

//this method computes the points for each receipt and stores in the map
func ProcessReceipt(receipt Receipt) (string, error, int) {

//initialize points variable
points := 0

//convert float from string 
total, err := strconv.ParseFloat(receipt.Total, 64)

//if value isn't float, return default error
if err != nil {
return "", fmt.Errorf(errMsg),badRequest
}

//compute points

//count number of alphanumeric characters in Retailer's name
  
for _, c := range receipt.Retailer{
if unicode.IsLetter(c) || unicode.IsDigit(c) {
points++
}
}

//check if round dollar 
if total == math.Floor(total){
points += 50
}

//check if multiple of 0.25
if math.Mod(total, 0.25) == 0 {
points += 25
}

//add points for every two items on the receipt
points += (5 * len(receipt.Items) / 2)

//examine trimmed length of each item's description 
for _, item := range receipt.Items {

//validate if description is present 
if len(item.ShortDescription) < 1{
return "", fmt.Errorf(errMsg), badRequest
}

//validate if price is a float value 
//extract price of current item
curPrice, err := strconv.ParseFloat(item.Price, 64)

//if price isn't a float value, return default error
if err != nil {
return "", fmt.Errorf(errMsg), badRequest
}

desc := strings.TrimSpace(item.ShortDescription)

//check if trimmed length is divisible by 3
if len(desc) % 3 == 0 {

points += int(math.Round(0.2 * curPrice))

}

}

//check purchase date
date := strings.Split(receipt.PurchaseDate, "-")

//if purchase date doesn't have 3 parts, then return error
if len(date) != 3 {
return "", fmt.Errorf(errMsg), badRequest
}

//purchase date is odd
purchDate, err := strconv.ParseInt(date[2], 10, 64)

if err != nil {
return "", fmt.Errorf(errMsg), badRequest
}

if purchDate % 2 != 0 {
points += 6
}

//check purchase time
time := strings.Split(receipt.PurchaseTime, ":")

//validate if purchase time has 2 parts with necessary delimiter
if len(time) != 2 {
return "", fmt.Errorf(errMsg), badRequest
}

//extract hour and minute
hour, err := strconv.ParseInt(time[0], 10, 8)

//if hour isn't an integer, return default error
if err != nil {
return "", fmt.Errorf(errMsg), badRequest
}

min, err := strconv.ParseInt(time[1], 10, 8)

//if minute isn't an integer, return default error
if err != nil {
return "", fmt.Errorf(errMsg), badRequest
}

//check if purchase time is between 2 and 4 
if ((hour == 14 && min > 0) || (hour > 14)) && (hour < 16) {
points += 10
}

id, err := generateId()

//if an error occurs while generating the id, return appropriate error message
if err != nil {
return "", fmt.Errorf("Internal Server Error. Please try again later."), internalServerError
}

//add current receipt to the map
ReceiptsMapper[id] = points

return id, nil, success
}

//function to generate unique id for receipt
func generateId()(string, error) {

id, err := exec.Command("uuidgen").Output()

//if an error was encountered while generating the ID, return it accordingly
if err != nil {
return "", fmt.Errorf("Internal Server Error. Please Try Again Later.")
}
//convert to string and lower case
idString := strings.ToLower(string(id))

//trim trailing and leading whitespaces
idString = strings.TrimSpace(idString)

return idString, nil
}

//retrieve points for given receipt id if valid, otherwise return error message

func RetrieveReceipts(id string) (int, error, int) {
id = strings.TrimSpace(id)
points, ok := ReceiptsMapper[id]

//if the id isn't valid, return error message accordingly
if !ok {
return 0, fmt.Errorf("No receipt found for that ID."), notFound
}
return points, nil, success

}
