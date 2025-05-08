package user

type User struct {
	ID    int    `json:"id,omitempty" validate:"omitempty"`
	Name  string `json:"name" validate:"required,min=2,max=50,notblank"`
	Email string `json:"email" validate:"required,email"`
	City  string `json:"city,omitempty" validate:"omitempty,min=2,max=100,notblank"`
	Age   int    `json:"age,omitempty" validate:"omitempty,min=2,max=100,notblank"`
}
