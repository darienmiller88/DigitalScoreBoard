package models

//struct to allow service to package data, error message and status code.
type Result[T any] struct{
	StatusCode int
	ResultData T
	Err        error
}