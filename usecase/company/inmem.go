package company

import (
	"TestVar/entity"
)

//inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.Company
}

//newInmem create new repository
func newInmem() *inmem {
	var m = map[entity.ID]*entity.Company{}
	return &inmem{
		m: m,
	}
}

//Create a Company
func (r *inmem) AddCompany(e *entity.Company) (entity.ID, error) {
	r.m[e.Id] = e
	return e.Id, nil
}

//Get a Company
func (r *inmem) Get(id entity.ID) (*entity.Company, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a Company
func (r *inmem) UpdateCompany(e *entity.Company) error {
	_, err := r.Get(e.Id)
	if err != nil {
		return err
	}
	r.m[e.Id] = e
	return nil
}

func (r *inmem) UpdateCompanyWithForm(e *entity.Company) error {
	_, err := r.Get(e.Id)
	if err != nil {
		return err
	}
	r.m[e.Id] = e
	return nil
}


/*//Search Companys
func (r *inmem) Search(query string) ([]*entity.Company, error) {
	var d []*entity.Company
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}
	return d, nil
}

//List Companys
func (r *inmem) List() ([]*entity.Company, error) {
	var d []*entity.Company
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}*/

//Delete a Company
func (r *inmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
