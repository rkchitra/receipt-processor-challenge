# Receipt Processor Challenge
Webservice for the fetch rewards receipt processor challenge

# Instructions to Run
1. Install the Mux dependency by running the following command:
   "go get -u github.com/gorilla/mux"

2. Compile and run all the files using the following command:
   "go run main.go handlers.go logic.go models.go"

# File Descriptions 
1. **main.go** : This file provides an entry point into the web service. The main function within it defines the end points and the corresponding handlers.
   
2. **handlers.go** : This file contains the handlers linked to the endpoints in the previous file (*main.go*). It calls the necessary functions to handle the underlying logic to service the requests. The functions within this file interpret the incoming request and also construct and send a JSON response accordingly.

3. **logic.go** : This file contains the functions that implement the required logic to service the API requests: they generate receipt Ids, compute the number of points allocated to each Id, and also store the mapping of receipt Ids to points for efficient retrieval.

4. **models.go** : This file defines the structs for JSON requests and responses, including error responses. 
