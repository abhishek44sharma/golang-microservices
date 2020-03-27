package domain

type User struct {
	Id			int64	`json:"id"`
	Firstname	string	`json:"first name"`
	Lastname	string	`json:"last_name"`
	Email		string	`json:"email"`
}