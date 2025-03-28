package models

// import(
// 	"github.com/go-ozzo/ozzo-validation"
// )

type Team struct{
	Score     int      `bson:"score"      json:"score"` 
	TeamName  string   `bson:"team_name"  json:"team_name"`
	Players   []string `bson:"players"    json:"players"`
}