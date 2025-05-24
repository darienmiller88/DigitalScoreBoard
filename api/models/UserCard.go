package models

//Temporarily removing ",omitempty"
type UserCard struct {
	Score int    `bson:"score" json:"score"` 
	Name  string `bson:"name"  json:"username"`
}

//Minified UserCard to send back to the client
type UserCardDTO struct{
	Name  string `json:"username"`
}