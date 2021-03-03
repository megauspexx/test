package handler

import (
	"TestVar/entity"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
	"TestVar/usecase/employee"
)

//import "net/http"

func AddEmployee(service employee.UseCase) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		errorMessage := "Error adding"
		var input struct {

			Name string `json:"name"`
			SecondName string `json:"secondName,omitempty"`
			Surname string `json:"surname,omitempty"`
			HireDate string `json:"hireDate,omitempty"`
			// employee position in the company
			Position string `json:"position,omitempty"`
			// ID of a company hired by
			CompanyId int64 `json:"companyId,omitempty"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateEmployee(input.Name, input.SecondName, input.Surname, input.HireDate, input.Position, input.CompanyId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &entity.Employee{
			Id:  id,
			Name :input.Name,
			SecondName:input.Name,
			Surname:input.Surname,
			HireDate:input.HireDate,
			// employee position in the company
			Position:input.Position,
			// ID of a company hired by
			CompanyId:input.CompanyId,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func DeleteEmployee(service employee.UseCase) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		errorMessage := "Error removing"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteEmployee(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func GetEmployeeById(service employee.UseCase) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		errorMessage := "Error reading"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetEmployeeById(id)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &entity.Employee{
			Id:       data.Id,
			Name:    data.Name,
			Surname:   data.Surname,
			CompanyId:    data.CompanyId,
			HireDate: data.HireDate,
			Position: data.Position,
			SecondName: data.SecondName,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}


func UpdateEmployeeWithForm(service employee.UseCase) http.Handler {


	 return Updateemployee(service)

	/*return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	})*/
}

func Updateemployee(service employee.UseCase) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		errorMessage := "Error adding"
		input :=  entity.Employee {}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		//id, err := service.UpdateEmployee(input.Name, input.SecondName, input.Surname, input.HireDate, input.Position, input.CompanyId)
		err = service.UpdateEmployee(&input)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &entity.Employee{
			Id:  input.Id,
			Name :input.Name,
			SecondName:input.Name,
			Surname:input.Surname,
			HireDate:input.HireDate,
			// employee position in the company
			Position:input.Position,
			// ID of a company hired by
			CompanyId:input.CompanyId,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func MakeEmployeeHandlers(r *mux.Router, n negroni.Negroni, service employee.UseCase) {

	r.Handle("/v2/employee", n.With(
		negroni.Wrap(UpdateEmployeeWithForm(service)),
	)).Methods("PUT", "OPTIONS").Name("updateEmployee")


	r.Handle("/v2/employee", n.With(
		negroni.Wrap(AddEmployee(service)),
	)).Methods("POST", "OPTIONS").Name("createEmployee")


	r.Handle("/v2/employee/{employeeId}", n.With(
		negroni.Wrap(Updateemployee(service)),
	)).Methods("POST", "OPTIONS").Name("createEmployee")


	r.Handle("/v2/employee/{employeeId}", n.With(
		negroni.Wrap(GetEmployeeById(service)),
	)).Methods("GET", "OPTIONS").Name("getEmployee")


	r.Handle("/v2/employee/{employeeId}", n.With(
		negroni.Wrap(DeleteEmployee(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteEmployee")
}


