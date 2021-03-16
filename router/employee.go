package router

import (
	controller "sql/controller"

	"github.com/gorilla/mux"
)

func RouterEmployee(myRouter *mux.Router) {
	myRouter.HandleFunc("/employee", controller.CreateEmployee).Methods("POST")
	myRouter.HandleFunc("/employees", controller.GetEmployee)
	myRouter.HandleFunc("/employee/{id}", controller.SingleEmployee)
	myRouter.HandleFunc("/employee", controller.UpdateEmployee).Methods("PUT")
	myRouter.HandleFunc("/employee", controller.DeleteEmployee).Methods("DELETE")
}
