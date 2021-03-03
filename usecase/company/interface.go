package company

import (
	"TestVar/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Company, error)
}

//Writer Company writer
type Writer interface {
	AddCompany(e *entity.Company) (entity.ID, error)
	UpdateCompanyWithForm(e *entity.Company) error
	UpdateCompany(e *entity.Company) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetCompanyById(id entity.ID) (*entity.Company, error)
	CreateCompany(	Name string, LegalForm string) (entity.ID, error)
	UpdateCompany(e *entity.Company) error
	DeleteCompany(id entity.ID) error
}
