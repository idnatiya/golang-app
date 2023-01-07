package types

type RegisterType struct {
	FirstName            string `json:"firstName" binding:"required"`
	LastName             string `json:"lastName" bindng:"required"`
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`
}
