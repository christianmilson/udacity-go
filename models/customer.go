package models

import (
	"math/rand"

	"github.com/bxcodec/faker/v4"
)

type Customer struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

func (customer *Customer) Faker() Customer {
	contacted := false
	randN := rand.Intn(2)
	if randN == 1 {
		contacted = true
	}

	return Customer{
		Id:        0,
		Name:      faker.FirstName(),
		Role:      faker.Word(),
		Email:     faker.Email(),
		Phone:     faker.Phonenumber(),
		Contacted: contacted,
	}
}
