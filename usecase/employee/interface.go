package employee

import (
	"TestVar/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Employee, error)
}

//Writer Employee writer
type Writer interface {
	AddEmployee(e *entity.Employee) (entity.ID, error)
	UpdateEmployeeWithForm(e *entity.Employee) error
	UpdateEmployee(e *entity.Employee) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetEmployeeById(id entity.ID) (*entity.Employee, error)
	CreateEmployee(	Name string, SecondName string, Surname string, HireDate string, Position string, CompanyId int64) (entity.ID, error)
	UpdateEmployee(e *entity.Employee) error
	DeleteEmployee(id entity.ID) error
}
