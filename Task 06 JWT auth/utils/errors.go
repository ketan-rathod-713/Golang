package utils

import (
	"encoding/json"
	"net/http"
	"task6MuxGorm/models"
)

/*Custom Response Errors */
var ERROR_DECODING_JSON = &models.ApiError{Code: http.StatusBadRequest, Message: "Error Decoding Json Data Provided"}
var ERROR_POST_REQ = &models.ApiError{Code: http.StatusBadRequest, Message: "Error Posting Data for given data"}
var ERROR_GET_REQ = &models.ApiError{Code: http.StatusBadRequest, Message: "Error Fetching Data"}
var ERROR_FINDING_ID = &models.ApiError{Code: http.StatusBadRequest, Message: "Error getting Id of type unsigned int"}
var ERROR_FINDING_COOKIE = &models.ApiError{Code: http.StatusBadRequest, Message: "Error getting cookie"}


// convert error into json and send it to client
func JSONError(w http.ResponseWriter, err *models.ApiError) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
}
