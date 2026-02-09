package entities

type UserEntity struct {
	Id       uint   `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
