package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

type ConnectionTestController struct{}

func (h *ConnectionTestController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			h.Get(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method Not Allowed")
	}
}

func (h *ConnectionTestController) Get(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pkg/views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error")
	}
	tmpl.Execute(w, nil)
}
