package models

//If the user logi with right credential a Token will be sent to the user
type JWT struct {
	//This field carries the value of the token
	Token string `json:"token"`
}
