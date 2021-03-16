package router

import (
	controller "sql/controller"

	"github.com/gorilla/mux"
)

func RouterTask(myRouter *mux.Router) {
	myRouter.HandleFunc("/task", controller.CreateTask).Methods("POST")
	myRouter.HandleFunc("/tasks", controller.GetTask)
	myRouter.HandleFunc("/task/{id}", controller.SingleTask)
	myRouter.HandleFunc("/task", controller.UpdateTask).Methods("PUT")
	myRouter.HandleFunc("/task", controller.DeleteTask).Methods("DELETE")
}
