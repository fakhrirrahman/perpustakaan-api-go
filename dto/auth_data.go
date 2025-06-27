package dto

// AuthData represents login credentials
type AuthData struct {
	Email    string `json:"email" validate:"required,email" example:"admin@example.com"`
	Password string `json:"password" validate:"required" example:"password123" minLength:"6"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}