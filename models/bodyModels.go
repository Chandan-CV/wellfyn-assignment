package models

// this is the model for the post request for signup (the body of the request)
type UserSignup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
