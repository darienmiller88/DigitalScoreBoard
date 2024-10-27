package models

type UserCard struct {
	Name      string `bson:"name"`
	Score     int    `bson:"score,omitempty"`
}