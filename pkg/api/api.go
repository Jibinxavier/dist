package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "TODO HEALTH\n")
}
func Status(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "TODO STATUS \n")
}
func RegisterWork(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "TODO RegisterWork \n")
}
func Job(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "TODO JOB \n")
}
func Assign(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "TODO ASSIGN JOB \n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

// create a
func NewRouter() {
	// factory  pattern
}
