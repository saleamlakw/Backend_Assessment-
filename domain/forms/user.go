package forms
type SignupForm struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=30"`
	LastName  string `json:"last_name" binding:"required,min=3,max=30"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=30,StrongPassword"`
}

type LoginForm struct {
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Password string `json:"password" bson:"password" binding:"required,min=4,max=30"`
}

type LoginResponseForm struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type RefreshTokenRequestForm struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}
type RefreshTokenResponseForm struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}