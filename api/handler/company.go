package handler

import (
	"TestVar/entity"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
	"TestVar/usecase/company"
)

//import "net/http"

func AddCompany(service company.UseCase) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		errorMessage := "Error adding"
		var input struct {

			Name string `json:"name,omitempty"`
			LegalForm string `json:"legalForm,omitempty"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateCompany(input.Name, input.LegalForm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &entity.Company{
			Id:  id,
			Name :input.Name,
			LegalForm:input.LegalForm,
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


/*func ListCompany(service company.UseCase) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		errorMessage := "Error List"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ, err  := service.ListCompanies()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}


		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}*/

func DeleteCompany(service company.UseCase) http.Handler {

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
		err = service.DeleteCompany(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func GetCompanyById(service company.UseCase) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		errorMessage := "Error reading"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["CompanyId"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.GetCompanyById(id)

		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		log.Println("ok1");
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &entity.Company{
			Id:       data.Id,
			Name:    data.Name,
			LegalForm:   data.LegalForm,
		}


		if err := json.NewEncoder(w).Encode(toJ); err != nil {

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func UpdateCompanyWithForm(service company.UseCase) http.Handler {


	return UpdateCompany(service)
}

func UpdateCompany(service company.UseCase) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		errorMessage := "Error adding"
		input :=  entity.Company {}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		//id, err := service.UpdateCompany(input.Name, input.SecondName, input.Surname, input.HireDate, input.Position, input.CompanyId)
		err = service.UpdateCompany(&input)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &entity.Company{
			Id:  input.Id,
			Name :input.Name,
			LegalForm:input.LegalForm,
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

func MakeCompanyHandlers(r *mux.Router, n negroni.Negroni, service company.UseCase) {

	r.Handle("/v2/company", n.With(
		negroni.Wrap(UpdateCompanyWithForm(service)),
	)).Methods("PUT", "OPTIONS").Name("updateCompany")


	r.Handle("/v2/company", n.With(
		negroni.Wrap(AddCompany(service)),
	)).Methods("POST", "OPTIONS").Name("createCompany")


	r.Handle("/v2/company/{CompanyId}", n.With(
		negroni.Wrap(UpdateCompany(service)),
	)).Methods("POST", "OPTIONS").Name("createCompany")


	r.Handle("/v2/company/{CompanyId}", n.With(
		negroni.Wrap(GetCompanyById(service)),
	)).Methods("GET", "OPTIONS").Name("getCompany")


	r.Handle("/v2/company/{CompanyId}", n.With(
		negroni.Wrap(DeleteCompany(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteCompany")



/*	r.Handle("/v2/company/{companyId}/employees", n.With(
		negroni.Wrap(ListCompany(service)),
	)).Methods("DELETE", "OPTIONS").Name("listCompany")*/
}


