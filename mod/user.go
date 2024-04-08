package mod

type User struct {
	UserId   string `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
