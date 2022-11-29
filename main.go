package main

import (
	"challange_shs_3/model"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDB() (*gorm.DB, error) {
	dbCredential := model.CredentialDB{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "",
		DatabaseName: "challange_database",
		Port:         "5432",
	}

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbCredential.Host, dbCredential.Username, dbCredential.Password, dbCredential.DatabaseName, dbCredential.Port)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	db, err := connectToDB()

	if err != nil {
		panic(err)
	}

	//DROP COLUMN IF EXIST
	db.Migrator().DropTable(&model.Employee{})
	db.Migrator().DropTable(&model.Departement{})

	//DO MIGRATION
	db.AutoMigrate(&model.Employee{}, &model.Departement{})

	//INSERT DEPARTEMENTS DATA
	var Departement = []model.Departement{
		{
			ID:              1,
			NamaDepartement: "FrontEnd",
		},
		{
			ID:              2,
			NamaDepartement: "BackEnd",
		},
	}

	db.Create(&Departement)

	//INSERT EMPLOYEES DATA
	var Employee = []model.Employee{
		{
			ID:            1,
			FirstName:     "Johan",
			LastName:      "Liebert",
			Email:         "johanliebert@gmail.com",
			City:          "Jakarta",
			DepartementID: 1,
		},
		{
			ID:            2,
			FirstName:     "Anna",
			LastName:      "Liebert",
			Email:         "annaliebert@gmail.com",
			City:          "Jakarta",
			DepartementID: 2,
		},
	}

	db.Create(&Employee)

	fmt.Print("Success!")

}
