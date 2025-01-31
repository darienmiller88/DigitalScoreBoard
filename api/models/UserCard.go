package models

//Temporarily removing ",omitempty"
type UserCard struct {
	Score     int    `bson:"score" json:"score"` 
	Name      string `bson:"name"  json:"username"`
}