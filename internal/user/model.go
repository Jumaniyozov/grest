package user

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	Email        string `json:"email" bson:"email"`
	PasswordHash string `json:"-" bson:"password"`
}

type CreateUserDTO struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
	Email        string `json:"email"`
}
