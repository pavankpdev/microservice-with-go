package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/util"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if reqBody.UserID == "" {
		http.Error(w, "Missing userID", http.StatusBadRequest)
		return
	}

	j, err := json.Marshal(reqBody)
	if err != nil {
		http.Error(w, "Failed to marshal request body", http.StatusInternalServerError)
		return
	}

	bytesBuffer := bytes.NewBuffer(j)

	resp, err := http.Post("http://trip-service:8083/preview", "application/json", bytesBuffer)
	if err != nil {
		http.Error(w, "Failed to call trip service", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var responseBody any
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		http.Error(w, "Invalid response body", http.StatusInternalServerError)
		return
	}

	resonse := contracts.APIResponse{
		Data:  responseBody,
		Error: nil,
	}

	util.WriteJSONResponse(w, http.StatusOK, resonse)
}
