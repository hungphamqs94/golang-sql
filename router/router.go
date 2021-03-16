package router

import (
	"fmt"
	"log"
	"net/http"
	"sql/controller"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", controller.GetAll)
	RouterEmployee(myRouter)
	RouterTask(myRouter)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
