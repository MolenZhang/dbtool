package main

import (
	"dbselector/controllers"
	"log"
	"net/http"
)

func main() {
	log.Println("main")
	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	http.HandleFunc("/", controllers.OpenHomePage)
	http.HandleFunc("/login", controllers.LoginAction)
	http.HandleFunc("/execsql", controllers.ExecSqlAction)
	http.HandleFunc("/exec", controllers.ExecAction)
	http.HandleFunc("/logout", controllers.LogoutAction)
	http.ListenAndServe(":8888", nil)
}
