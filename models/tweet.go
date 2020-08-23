package models

// Tweet is tweet message
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
