package models

type LocationResult struct{
	StatusCode int
	Location   Location
	Err        error
}