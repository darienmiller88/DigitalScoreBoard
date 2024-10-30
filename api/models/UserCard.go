package models

type UserCard struct {
	Score     int    `bson:"score,omitempty" json:"score"` 
	Name      string `bson:"name"            json:"name"`
}