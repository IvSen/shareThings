package dto

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
	Country  string
	City     string
	District string
	Postcode string
}

type UpdateUserRequest struct {
	Name     string
	Email    string
	Country  string
	City     string
	District string
	Postcode string
}

type SignInUserDTO struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenRequest struct {
	Token string `json:"token" binding:"required"`
}

type GetUser struct {
	Id string
}
