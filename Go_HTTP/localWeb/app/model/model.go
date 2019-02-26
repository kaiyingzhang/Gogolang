package model

import "github.com/jinzhu/gorm"

type Employee struct {
	// gorm.Model
	Id   int    `json:"id"`
	Name string `json:"name"`
	Company int `gorm:"ForeignKey:comid" json:"company"`
}

type Employees []Employee

type Company struct {
	gorm.Model
	comId int `json: "comid"`
	comName string `json : "comname"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Employee{}, &Company{})
	db.Model(&Company{}).AddForeignKey("comid", "Employee(id)", "CASCADE", "CASCADE")
	return db
}
