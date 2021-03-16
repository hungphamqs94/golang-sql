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

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	sqlStatement := "INSERT INTO employee(name,username,password,role,email,phone) values(?,?,?,?,?,?)"
	reqBody, _ := ioutil.ReadAll(r.Body)
	var employee models.Employee
	json.Unmarshal(reqBody, &employee)
	_, errInsert := singleton.GetInstance().Exec(sqlStatement, employee.Name, employee.Username, employee.Password, employee.Role, employee.Email, employee.Phone)
	if errInsert != nil {
		panic(errInsert.Error())
	}
	json.NewEncoder(w).Encode(employee)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	results, _ := singleton.GetInstance().Query("Select * from Employee")
	var listEmployee []models.Employee
	for results.Next() {
		var employee models.Employee
		errEmployee := results.Scan(&employee.Id, &employee.Name, &employee.Username, &employee.Password, &employee.Role, &employee.Email, &employee.Phone)
		if errEmployee != nil {
			panic(errEmployee.Error())
		}
		listEmployee = append(listEmployee, employee)
	}
	json.NewEncoder(w).Encode(listEmployee)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	sqlStatement := "update Employee set username=? where id=?"
	reqBody, _ := ioutil.ReadAll(r.Body)
	var employee models.Employee
	json.Unmarshal(reqBody, &employee)
	_, errUpdate := singleton.GetInstance().Exec(sqlStatement, employee.Username, employee.Id)
	if errUpdate != nil {
		panic(errUpdate.Error())
	}
	json.NewEncoder(w).Encode("Update Employee thành công")
}

func ConCurrencyEmployee(chanEmployee chan []models.Employee) {
	sqlStatementEmployee := "Select * from employee"
	var listEmployee []models.Employee
	cursor, err := singleton.GetInstance().Query(sqlStatementEmployee)
	if err != nil {
		panic(err.Error())
	}
	for cursor.Next() {
		var employee models.Employee
		employeeErr := cursor.Scan(&employee.Id, &employee.Name, &employee.Username, &employee.Password, &employee.Role, &employee.Email, &employee.Phone)
		if employeeErr != nil {
			panic(employeeErr.Error())
		}
		listEmployee = append(listEmployee, employee)
	}
	chanEmployee <- listEmployee
}

func ConCurrencyTask(chanTask chan []models.Task) {
	sqlStatementTask := "Select * from task"
	var listTask []models.Task
	cursor, err := singleton.GetInstance().Query(sqlStatementTask)
	if err != nil {
		panic(err.Error())
	}
	for cursor.Next() {
		var task models.Task
		taskErr := cursor.Scan(&task.Id, &task.Name, &task.AssignedId, &task.Prioty, &task.Hour)
		if taskErr != nil {
			panic(taskErr.Error())
		}
		listTask = append(listTask, task)
	}
	chanTask <- listTask
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	var chanEmploye chan []models.Employee
	var chanTask chan []models.Task
	go ConCurrencyEmployee(chanEmploye)
	go ConCurrencyTask(chanTask)
	listEmployee := <-chanEmploye
	listTask := <-chanTask
	json.NewEncoder(w).Encode(listEmployee)
	json.NewEncoder(w).Encode(listTask)

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	sqlStatement := "delete from Employee where id=?"
	reqBody, _ := ioutil.ReadAll(r.Body)
	var id int
	json.Unmarshal(reqBody, &id)
	_, errDelete := singleton.GetInstance().Exec(sqlStatement, id)
	if errDelete != nil {
		panic(errDelete.Error())
	}
	json.NewEncoder(w).Encode("Delete Employee thành công")
}

func SingleEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	sqlStatement := "select * from Employee where id=?"
	result, _ := singleton.GetInstance().Query(sqlStatement, key)
	for result.Next() {
		var employee models.Employee
		errEmployee := result.Scan(&employee.Id, &employee.Name, &employee.Username, &employee.Password, &employee.Role, &employee.Email, &employee.Phone)
		if errEmployee != nil {
			panic(errEmployee.Error())
		}
		json.NewEncoder(w).Encode(employee)
	}
}
