package server

import (
	"fmt"
	"net/http"
)

func (mux *HttpI) router() {
	mux.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Olá mundo!")
	})
}
