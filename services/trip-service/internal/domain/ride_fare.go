package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	ID                primitive.ObjectID `json:"id"`
	UserID            string             `json:"userID"`
	PackageSlug       string             `json:"packageSlug"`
	TotalPriceInCents float64            `json:"totalPriceInCents"`
}
