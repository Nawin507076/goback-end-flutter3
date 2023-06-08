package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	PhoneNumber string    `json:"phoneNumber"`
	Ages        string    `json:"ages"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"createdAt"`
}

func NewAccount(id int, firstName string, lastName string, phoneNumber string, ages string, email string) *Account {
	return &Account{
		ID:          rand.Intn(10000),
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Ages:        ages,
		Email:       email,
		CreatedAt:   time.Now().UTC(),
	}
}
