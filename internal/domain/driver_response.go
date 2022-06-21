package domain

type DriverResponse struct {
	ID       int     `db:"id" json:"id" binding:"omitempty"`
	Name     string  `db:"name" json:"name" binding:"omitempty"`
	Phone    string  `db:"phone" json:"phone" binding:"required,e164"`
	Email    string  `db:"email" json:"email" binding:"required,email"`
	Rating   float32 `db:"rating" json:"rating" binding:"omitempty"`
	TaxiType string  `db:"taxi_type" json:"taxiType" binding:"required"`
}
