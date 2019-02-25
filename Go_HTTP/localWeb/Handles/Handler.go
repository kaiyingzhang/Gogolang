package Handles

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"net/http"

	"../Models/Employees.go"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request)  {
	todos := Employees_go.Employees{
		Employee{Name:"Write presentation"},
		Employee{Name:"Host meetup"},
		Employee{Id:22,Name:"Joel"},
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
