package models

// LogInResponse stores token used in LogIn function
type LogInResponse struct {
	Token string `json:"token,omitempty"`
}
