package models

import "time"

type Address struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type DbSettings struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Note struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Active    bool      `json:"active"`
}

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
