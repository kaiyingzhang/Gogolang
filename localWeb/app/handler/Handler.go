package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kaiyingzhang/Gogolang/localWeb/model"
	"net/http"
)


func init() {
	fmt.Println("handler package initialized")
}

func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Welcome!")
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request)  {
	todos := model.Employees{
		model.Employee{Name:"Write presentation"},
		model.Employee{Name:"Host meetup"},
		model.Employee{Id:01,Name:"Kay Zhang"},
		model.Employee{Id:02,Name:"Lesli Xue"},
	}
	if err := json.NewEncoder(w).Encode(todos); err != nil{
		panic(err)
	}
}


func TodoShow(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
