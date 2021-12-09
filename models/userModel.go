package models

type User struct {
	Id        int64   `json:"id,omitempty"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Phone     string  `json:"phone"`
	Address   string  `json:"address"`
	City      string  `json:"city"`
	UserType  string  `json:"userType"`
	Lat       float32 `json:"lat"`
	Lng       float32 `json:"lng"`
	Status    int     `json:"status"`
	Created   string  `json:"created"`
}
