package model

import "time"

type User struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Session struct {
	ID     int       `json:"id"`
	Token  string    `json:"token"`
	Name   string    `json:"name"`
	Expiry time.Time `json:"expiry"`
}

type Book struct {
	ID     int    `json:"book_id"`
	Code   string `json:"code"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  int    `json:"stock"`
	Status string `json:"-"`
}

type Borrowed struct {
	ID         int       `json:"id"`
	CodeBook   string    `json:"code_book"`   //berisi code buku
	BorrowedBy int       `json:"borrowed_by"` //berisi code member
	BorrowDate time.Time `json:"borrow_date"`
	ReturnDate time.Time `json:"return_date"`
	Fine       int       `json:"fine"` //denda
}

type Return struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
