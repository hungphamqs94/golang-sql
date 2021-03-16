package models

type Employee struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}