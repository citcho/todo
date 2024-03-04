package user

func NewUserFixture(id string, name string, email string, password string) *User {
	return &User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}
}
