package http

import (
	"encoding/json"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
	"ride-sharing/shared/util"
)

type HttpHandler struct {
	Service domain.TripService
}

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (h *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
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

	ctx := r.Context()
	fare := &domain.RideFareModel{
		UserID:            reqBody.UserID,
		PackageSlug:       "standard",
		TotalPriceInCents: 1000,
	}

	t, err := h.Service.CreateTrip(ctx, fare)
	if err != nil {
		http.Error(w, "Failed to create trip", http.StatusInternalServerError)
		return
	}

	util.WriteJSONResponse(w, http.StatusOK, t)
}
