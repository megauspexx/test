package entity

type Employee struct {

	Id ID `json:"id,omitempty"`

	Name string `json:"name"`

	SecondName string `json:"secondName,omitempty"`

	Surname string `json:"surname,omitempty"`

	HireDate string `json:"hireDate,omitempty"`

	// employee position in the company
	Position string `json:"position,omitempty"`

	// ID of a company hired by
	CompanyId int64 `json:"companyId,omitempty"`
}


func NewEmployee(name string, secondName string, surname string, hireDate string, position string, companyId int64) (*Employee, error) {

	b := &Employee{
		Id:        NewID(),
		Name:     name,
		SecondName:    secondName,
		Surname:     surname,
		HireDate:  hireDate,
		Position: position,
		CompanyId: companyId,
	}
	err := b.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return b, nil
}

//Validate validate
func (e *Employee) Validate() error {
	if e.Name == ""  {
		return ErrInvalidEntity
	}
	return nil
}

