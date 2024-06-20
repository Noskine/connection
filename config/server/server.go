package server

import (
	"fmt"
	"log"
	"net/http"
)

type HttpI struct {
	mux *http.ServeMux
}

func NewHttpI() *HttpI {
	return &HttpI{
		mux: http.NewServeMux(),
	}
}

func (ht *HttpI) Listen(port int, message string) {
	ht.router()

	Sport := fmt.Sprintf(":%v", port)
	logger :=  fmt.Sprintf("%s :: ==> http://localhost%v", message, Sport)

	log.Println(logger)

	err := http.ListenAndServe(Sport, ht.mux)
	if err != nil {
		panic(err.Error())
	}
}