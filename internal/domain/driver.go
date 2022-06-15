package domain

type Driver struct {
	ID           int    `db:"id" json:"id" binding:"omitempty"`
	Name         string `db:"name" json:"name" binding:"omitempty"`
	Phone        string `db:"phone" json:"phone" binding:"required,e164"`
	PasswordHash string `db:"password_hash" json:"passwordHash" binding:"required"`
	Email        string `db:"email" json:"email" binding:"required,email"`
}
