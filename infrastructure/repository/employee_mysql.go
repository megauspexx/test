package repository

import (
	"TestVar/entity"
	"database/sql"
	"log"
)

//mysql repo
type EmployeeMysql struct {
	db *sql.DB
}

//create new repository
func NewEmployeeMysql(db *sql.DB) *EmployeeMysql {
	return &EmployeeMysql{
		db: db,
	}
}

//Create
func (r *EmployeeMysql) AddEmployee(e *entity.Employee) (entity.ID, error) {

	e.Id = entity.NewID()
	log.Print("Addemployee")
	return e.Id, nil
}

//Get
func (r *EmployeeMysql) Get(id entity.ID) (*entity.Employee, error) {

	b := entity.Employee{Id:id}

	return &b, nil
}

//Update
func (r *EmployeeMysql) UpdateEmployeeWithForm(e *entity.Employee) error {

	log.Print("UpdateEmployeeWithForm with id:", e.Id)
	return nil
}

//Update
func (r *EmployeeMysql) UpdateEmployee(e *entity.Employee) error {

	log.Print("UpdateEmployee with id:", e.Id)
	return nil
}

//Delete
func (r *EmployeeMysql) Delete(id entity.ID) error {

	log.Print("DeleteEmployee with id:", id)
	return nil
}
