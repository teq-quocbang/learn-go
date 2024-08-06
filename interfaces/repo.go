package interfaces

type User struct {
	Email    string
	Name     string
	Provider string
}

// method with interface
// abstract
type Repository interface {
	GetUserInfo(id int) (User, error)
}

type AuthRepository interface {
	Login() error
}
