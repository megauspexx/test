package main

import (
	"TestVar/api/handler"
	"TestVar/api/middleware"
	"TestVar/infrastructure/repository"
	"TestVar/usecase/company"
	"TestVar/usecase/employee"
	"database/sql"
	"fmt"
	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load("./config.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	/*
		listener, err := net.Listen("tcp", ":5300")
		if err != nil {
			grpclog.Fatalf("failed to listen: %v", err)
		}
	*/


	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", os.Getenv("DB_USER") , os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	employeeRepo := repository.NewEmployeeMysql(db)
	employeeService := employee.NewService(employeeRepo)

	/*server := grpc.NewServer()
	pb.RegisterEmployeeServer(server, employeeService)

	err = server.Serve(list)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}*/



	companyRepo := repository.NewCompanyMysql(db)
	companyService := company.NewService(companyRepo)

	//loanUseCase := loan.NewService(userService, employeeService)

	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	handler.MakeEmployeeHandlers(r, *n, employeeService)


	handler.MakeCompanyHandlers(r, *n, companyService)


	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + os.Getenv("API_PORT"),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
