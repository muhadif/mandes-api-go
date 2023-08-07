package entity

type LoginRequest struct {
	Email		string
	Password	string
}

type LoginResponse struct {
	Token		string
	RefreshToken string
}

type RegisterRequest struct {
	Email		string
	FullName	string
	Password	string
}

type RegisterFromAdminRequest struct {
	Email		string
	FullName	string
	Password	string
	Role		UserRole
}
