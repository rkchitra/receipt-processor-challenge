/*
this file defines the necessary custom structures defined that are used to
interpret JSON requests and format JSON responses
*/
package main

//defines the structure of each item within a receipt
type Item struct {
ShortDescription string `json:"shortDescription"`
Price string `json:"price"`
}

//defines the structure of each item in the receipt payload
type Receipt struct {
Retailer string `json:"retailer"`
PurchaseDate string `json:"purchaseDate"`
PurchaseTime string `json:"purchaseTime"`
Items []Item `json:"items"`
Total string `json:"total"`
}

/*
defines the structure of the JSON response for POST method 
(receipt is presented)
*/
type ReceiptResponse struct {
Id string `json:"id"`
}

/*
defines the structure of the JSON response while returning the number of points
for valid receipt ID
*/
type PointsResponse struct {
Points int `json:points`
}

/*
defines the structure of the JSON response when there's an error
*/
type ErrorResponse struct {
Error string `json:"error"`
}
