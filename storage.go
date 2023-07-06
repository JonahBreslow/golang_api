package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(a *Account) error
	DeleteAccount(int) error
	UpdateAccount(a *Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print("Error: " + err.Error())
		return nil, err
	}
	if err := db.Ping(); err != nil {
		fmt.Print("Ping Error: " + err.Error())
		return nil, err
	}
	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStorage) createAccountTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY, 
		first_name TEXT, 
		last_name TEXT, 
		number BIGINT, 
		balance BIGINT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateAccount(a *Account) error {
	query := `INSERT INTO account 
	(first_name, last_name, number, balance, created_at) 
	VALUES ($1, $2, $3, $4, $5) `

	resp, err := s.db.Query(
		query, 
		a.FirstName, 
		a.LastName, 
		a.Number, 
		a.Balance, 
		a.CreatedAt,)
	if err != nil {
		fmt.Print("Error: " + err.Error())
		return err
	}
	
	fmt.Printf("Response: %v", resp)
	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStorage) UpdateAccount(a *Account) error {
	return nil
}

func (s *PostgresStorage) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStorage) GetAccounts() ([]*Account, error) {

	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.Id, 
			&account.FirstName, 
			&account.LastName, 
			&account.Number, 
			&account.Balance, 
			&account.CreatedAt)
		
		if err != nil {
			return nil, err
		}
	
		accounts = append(accounts, account)
	}

	return accounts, nil
}