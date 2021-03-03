package employee

import (
	"TestVar/entity"
)

//inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.Employee
}

//newInmem create new repository
func newInmem() *inmem {
	var m = map[entity.ID]*entity.Employee{}
	return &inmem{
		m: m,
	}
}

//Create a Employee
func (r *inmem) AddEmployee(e *entity.Employee) (entity.ID, error) {
	r.m[e.Id] = e
	return e.Id, nil
}

//Get a Employee
func (r *inmem) Get(id entity.ID) (*entity.Employee, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a Employee
func (r *inmem) UpdateEmployee(e *entity.Employee) error {
	_, err := r.Get(e.Id)
	if err != nil {
		return err
	}
	r.m[e.Id] = e
	return nil
}

//Update a Employee
func (r *inmem) UpdateEmployeeWithForm(e *entity.Employee) error {
	_, err := r.Get(e.Id)
	if err != nil {
		return err
	}
	r.m[e.Id] = e
	return nil
}

/*//Search Employees
func (r *inmem) Search(query string) ([]*entity.Employee, error) {
	var d []*entity.Employee
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}
	return d, nil
}

//List Employees
func (r *inmem) List() ([]*entity.Employee, error) {
	var d []*entity.Employee
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}*/

//Delete a Employee
func (r *inmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
