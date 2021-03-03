package repository

import (
	"TestVar/entity"
	"database/sql"
	"log"
)

//mysql repo
type CompanyMysql struct {
	db *sql.DB
}

//create new repository
func NewCompanyMysql(db *sql.DB) *CompanyMysql {
	return &CompanyMysql{
		db: db,
	}
}

//Create
func (r *CompanyMysql) AddCompany(e *entity.Company) (entity.ID, error) {

	e.Id = entity.NewID()
	log.Print("AddCompany")
	return e.Id, nil
}

//Get
func (r *CompanyMysql) Get(id entity.ID) (*entity.Company, error) {

	b := entity.Company{Id:id}
	return &b, nil
}

//Update
func (r *CompanyMysql) UpdateCompanyWithForm(e *entity.Company) error {

	log.Print("UpdateCompanyWithForm with id:", e.Id)
	return nil
}

//Update
func (r *CompanyMysql) UpdateCompany(e *entity.Company) error {

	log.Print("UpdateCompany with id:", e.Id)
	return nil
}

//Delete
func (r *CompanyMysql) Delete(id entity.ID) error {

	log.Print("DeleteCompany with id:", id)
	return nil
}