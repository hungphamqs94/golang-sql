package models

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	AssignedId int `json:"assignedId"`
	Prioty     int `json:"prioty"`
	Hour       int `json:"hour"`
}