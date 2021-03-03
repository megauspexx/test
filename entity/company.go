package entity

type Company struct {

	Id ID `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	// Company legal form (OOO, ZAO, IP)
	LegalForm string `json:"legalForm,omitempty"`
}


func NewCompany(name string, legalForm string) (*Company, error) {

	b := &Company{
		Id:        NewID(),
		Name:     name,
		LegalForm:legalForm,
	}

	err := b.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return b, nil
}

//Validate validate
func (e *Company) Validate() error {
	if e.Name == "" || e.LegalForm == ""  {
		return ErrInvalidEntity
	}
	return nil
}
