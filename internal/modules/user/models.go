package user

type User struct {
	ID        int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"omitempty,min=2,max=50"`
	LastName  string `json:"last_name" validate:"omitempty,min=2,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Age       int    `json:"age" validate:"omitempty,gte=5,lte=120"`
	Password  string `json:"password" validate:"required,min=8"`
}
