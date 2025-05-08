package userdto

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=50,notblank"`
	Email string `json:"email,omitempty" validate:"omitempty,email" example:"yatender@gmail.com"`
	City  string `json:"city,omitempty" validate:"omitempty,min=2,max=100,notblank"`
	Age   int    `json:"age,omitempty" validate:"omitempty,min=2,max=100,notblank" example:"30"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" omitempty:"required,min=2,max=50,notblank"`
	Email string `json:"email,omitempty" validate:"omitempty,email" example:"yatender@gmail.com"`
	City  string `json:"city,omitempty" validate:"required,min=2,max=100,notblank"`
	Age   int    `json:"age,omitempty" validate:"omitempty,min=2,max=100,notblank" example:"30"`
}
