package employee

import (
	"TestVar/entity"
)

//Service employee usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Create
func (s *Service) CreateEmployee(name string, secondName string, surname string, hireDate string, position string, companyId int64) (entity.ID, error) {
	b, err := entity.NewEmployee(name, secondName, surname , hireDate, position, companyId)
	if err != nil {
		return b.Id, err
	}
	return s.repo.AddEmployee(b)
}

//Get
func (s *Service) GetEmployeeById(id entity.ID) (*entity.Employee, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}


//Delete
func (s *Service) DeleteEmployee(id entity.ID) error {
	_, err := s.GetEmployeeById(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//Update
func (s *Service) UpdateEmployee(e *entity.Employee) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	//e.UpdatedAt = time.Now()
	return s.repo.UpdateEmployee(e)
}
