package model

import "time"

type User struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type MemberBorrow struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	BorrowCount int    `json:"borrow_count"`
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
	ID           int       `json:"id"`
	CodeBook     string    `json:"code_book"`   //berisi code buku
	CodeMember   string    `json:"code_member"` //berisi code member
	BorrowedDate time.Time `json:"borrow_date"`
	ReturnedDate time.Time `json:"return_date"`
	Late         int       `json:"late"`
	Quantity     int       `json:"quantity"`
	Status       string    `json:"status"` // borrowed/Not Borrowed
}

type Penalties struct {
	ID            int       `json:"id"`
	CodeMember    string    `json:"code_member"`
	PenaltyType   string    `json:"pinalty_type"`
	PenaltyAmount float32   `json:"penalty_amount"`
	PenaltyDate   time.Time `json:"penalty_date"`
	ResolveDate   time.Time `json:"resolve_date"`
	PenaltyActive bool      `json:"pinalty_active"`
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
