package main

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {

	ctx := context.Background()
	inmemRepo := repository.NewInMemoryTripRepository()
	service := service.NewService(inmemRepo)

	fare := &domain.RideFareModel{
		ID:                primitive.NewObjectID(),
		UserID:            "user123",
		PackageSlug:       "standard",
		TotalPriceInCents: 1500,
	}

	service.CreateTrip(ctx, fare)
	println("Trip service is running with in-memory repository:", service)
}
