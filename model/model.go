package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	ID            int `gorm:"primaryKey"`
	FirstName     string
	LastName      string
	Email         string
	City          string
	DepartementID int
	Departement   Departement
	gorm.Model
}

type Departement struct {
	ID              int `gorm:"primaryKey"`
	NamaDepartement string
	gorm.Model
}

type CredentialDB struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         string
	Schema       string
}
