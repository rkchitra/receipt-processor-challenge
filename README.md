# Receipt Processor Challenge
Webservice for the fetch rewards receipt processor challenge.

# Instructions to Run
1. Install the Mux dependency by running the following command: <br>
   "go get -u github.com/gorilla/mux"

2. Compile and run all the files using the following command: <br>
   "go run main.go handlers.go logic.go models.go"

# File Descriptions 
1. **main.go** : This file provides an entry point into the web service. The main function within it defines the end points and the corresponding handlers.
   
2. **handlers.go** : This file contains the handlers linked to the endpoints in the previous file (*main.go*). It calls the necessary functions to handle the underlying logic to service the requests. The functions within this file interpret the incoming request and also construct and send a JSON response accordingly. It also performs some validations on the JSON request (for the POST method). 

3. **logic.go** : This file contains the functions that implement the required logic to service the API requests: they generate receipt Ids, compute the number of points allocated to each Id, and also store the mapping of receipt Ids to points for efficient retrieval. It also performs certain validations to ensure the fields are of the expected data type. 

4. **models.go** : This file defines the structs for JSON requests and responses, including error responses. 

# Testing 
1. **POST Method** <br>
   URL: http://localhost:8080/receipts/process

   Sample JSON payloads:
   
```json
   {
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```
Total Points: 28


```json
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
```
Total Points: 109

2. **GET Method** <br>
   URL: http://localhost:8080/receipts/{id}/points <br>
   For id, plug in the id obtained in the JSON response of a POST method in the same session.
