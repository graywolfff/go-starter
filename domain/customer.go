package domain

import "time"

type Customer struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Zipcode     string    `json:"zipcode"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Status      string    `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]*Customer, error)
}
