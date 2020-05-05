package models

//For sending the error to the client
type Error struct {
	Message string `json:"message"`
}
