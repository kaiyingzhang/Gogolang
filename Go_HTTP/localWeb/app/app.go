package app

import (
	"./handler"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)
// App has router and db instances
type App struct {
	Router      *mux.Router
	DB     *gorm.DB
}

// // Initialize initializes the app with predefined configuration
// func (a *App) Initialize(config *config.Config) {
// 	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
// 		config.DB.Username,
// 		config.DB.Password,
// 		config.DB.Name,
// 		config.DB.Charset)
//
// 	db, err := gorm.Open(config.DB.Dialect, dbURI)
//
// 	if err != nil {
// 		log.Fatal("Could not connect database")
// 	}
//
// 	a.DB = model.DBMigrate(db)
//
// 	a.Router = mux.NewRouter()
// 	a.setRouters()
// }

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}
func (a *App) setRouters() {
	a.Get("/api/employee", a.GetAllEmployees)
	a.Get("/",a.GetInitial)
}


// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}
/*
** Projects Handlers
 */
func (a *App) GetInitial(w http.ResponseWriter, r *http.Request) {
	handler.Index(w, r)
}

func (a *App) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handler.GetAllEmployees(w, r)
}


func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}