package mod

type User struct {
	UserId   int64  `db:"user_id,string"`
	Username string `db:"username"`
	Password string `db:"password"`
	Token    string
}
