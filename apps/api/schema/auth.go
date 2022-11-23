package schema

type Register struct {
	Username string `validate:"required,min=2,max=15" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6" json:"password"`
}

type Login struct {
	Username string `validate:"required,min=2,max=15" json:"username"`
	Password string `validate:"required,min=6" json:"password"`
}
