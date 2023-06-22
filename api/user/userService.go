package user

type UserService interface {
	CreateUser(*User) error
	FindUserByEmail(string) (*User, error)
}
