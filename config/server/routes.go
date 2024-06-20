package server

import (
	"fmt"
	"net/http"

	"github.com/Noskine/connection/internal/controllers"
)

func (mux *HttpI) router() {
	mux.loadingFileStatic()

	mux.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Olá mundo!")
	})

	mux.mux.Handle("/helloworld", &controllers.ConnectionTestController{})
}
