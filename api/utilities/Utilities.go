package utilities

import (
	"encoding/json"
	"net/http"
)

//Utility function to allow the sending of json back to the client.
func SendJSON[T any](statusCode int, res http.ResponseWriter, data T) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(data)
}