package user

func NewUserFixture(ulid string, name string, email string, password string) *User {
	return &User{
		ulid:     ulid,
		name:     name,
		email:    email,
		password: password,
	}
}
