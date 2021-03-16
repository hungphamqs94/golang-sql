package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	models "sql/models"

	singleton "sql/singleton"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var task models.Task
	json.Unmarshal(reqBody, &task)

	sqlStatement := "INSERT INTO Task(name, assignedId, prioty, hour) values(?,?,?,?)"
	_, err := singleton.GetInstance().Exec(sqlStatement, task.Name, task.AssignedId, task.Prioty, task.Hour)

	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(task)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	sqlStatement := "Select * from task"
	results, errQuery := singleton.GetInstance().Query(sqlStatement)
	if errQuery != nil {
		panic(errQuery.Error())
	}

	var tasks []models.Task

	for results.Next() {
		var task models.Task
		errScan := results.Scan(&task.Id, &task.Name, &task.AssignedId, &task.Prioty, &task.Hour)
		if errScan != nil {
			panic(errScan.Error())
		}
		tasks = append(tasks, task)
	}
	json.NewEncoder(w).Encode(tasks)
	results.Close()
}

func SingleTask(w http.ResponseWriter, r *http.Request) {
	sqlStatement := "Select * from task where id=?"
	vars := mux.Vars(r)
	key := vars["id"]
	result, errQuery := singleton.GetInstance().Query(sqlStatement, key)
	if errQuery != nil {
		panic(errQuery.Error())
	}
	var task models.Task
	for result.Next() {
		errScan := result.Scan(&task.Id, &task.Name, &task.AssignedId, &task.Prioty, &task.Hour)
		if errScan != nil {
			panic(errScan.Error())
		}
		json.NewEncoder(w).Encode(task)
	}
	result.Close()
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	sqlStatement := "update task set name='Lam App  di dong' where id=?"
	reqBody, _ := ioutil.ReadAll(r.Body)
	var id int
	json.Unmarshal(reqBody, &id)
	_, errUpdate := singleton.GetInstance().Exec(sqlStatement, id)
	if errUpdate != nil {
		panic(errUpdate.Error())
	}
	json.NewEncoder(w).Encode("Update thành công")
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	sqlStatement := "Delete from task where id=?"
	reqBody, _ := ioutil.ReadAll(r.Body)
	var id int
	json.Unmarshal(reqBody, &id)
	_, errDelete := singleton.GetInstance().Exec(sqlStatement, id)
	if errDelete != nil {
		panic(errDelete.Error())
	}
	json.NewEncoder(w).Encode("Delete thành công")
}
