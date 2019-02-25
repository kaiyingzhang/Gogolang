package dbs

import "net/http"

type Route struct {
	Id int
	Name string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{Name:"Index",     Method:"GET",   Pattern:"/",                HandlerFunc:Handler.Index},
	Route{Name:"TodoIndex", Method:"GET",   Pattern:"/apis",           HandlerFunc:Handler_go.TodoIndex},
	Route{Name:"TodoShow",  Method:"GET",   Pattern:"/apis/{getEmployee}",  HandlerFunc:Handler_go.TodoShow},
}


