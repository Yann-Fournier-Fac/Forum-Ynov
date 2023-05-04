package database

type userInfo struct {
	Email    string `sql:"email"`
	Username string `sql:"username"`
	Password string `sql:"password"`
}
