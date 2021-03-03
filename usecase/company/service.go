package company

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

//
func (s *Service) CreateCompany(name string, legalForm string) (entity.ID, error) {
	c, err := entity.NewCompany(name, legalForm)
	if err != nil {
		return c.Id, err
	}
	return s.repo.AddCompany(c)
}

//Get
func (s *Service) GetCompanyById(id entity.ID) (*entity.Company, error) {
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
func (s *Service) DeleteCompany(id entity.ID) error {
	_, err := s.GetCompanyById(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//Update
func (s *Service) UpdateCompany(e *entity.Company) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	return s.repo.UpdateCompany(e)
}
