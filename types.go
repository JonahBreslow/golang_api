package main

import "math/rand"

type Account struct {
	Id       int    `json:"id"`
	FirsName string `json:"firstName"`
	LastName string `json:"lastName"`
	Number   int64  `json:"number"`
	Balance  int64  `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		Id:       rand.Intn(100000),
		FirsName: firstName,
		LastName: lastName,
		Number:   int64(rand.Intn(1000000000000000)),
	}
}
