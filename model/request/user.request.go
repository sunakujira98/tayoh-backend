package request

type UserCreateRequest struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required, email"`
	Address	string `json:"address"`
	Phone	string `json:"phone"`
	Password string `json:"password" validate:"required, min=6"`
}