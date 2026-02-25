package dtos

type SignUpRequest struct {
	Username string `json:"username" binding:"required" example:"qaseh_dev"`
	Email    string `json:"email" binding:"required,email" example:"qaseh@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"secret123"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"qaseh_dev"`
	Password string `json:"password" binding:"required" example:"secret123"`
}
