package auth

type AuthRequestDTO struct {
	Account  string `json:"account" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6"`
}
