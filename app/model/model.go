package model

import (
	"github.com/jinzhu/gorm"
	//import gorm mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Employee Struct
type Employee struct {
	gorm.Model
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Age    int    `json:"age"`
	Status bool   `json:"status"`
}

//Disable Employee
func (e *Employee) Disable() {
	e.Status = false
}

//Enable Employee
func (e *Employee) Enable() {
	e.Status = true
}

// DBMigrate will create and migrate the tables
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Employee{})
	return db
}
