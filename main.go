package main

import (
  "log"
	"github.com/bmizerany/pat"
	"net/http"
	"runtime"

	"github.com/trevrosen/target-snitch/platform/generic"
	"github.com/trevrosen/target-snitch/platform/osx"
)



func main() {
	m := pat.New()

	var informant *generic.GenericInformant
	switch runtime.GOOS {
	case "darwin":
		informant = osx.New(m)
	}


	loadMuxedRoutes(informant)
	http.Handle("/", m)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func loadMuxedRoutes(informant *generic.GenericInformant){
	informant.Pat.Get("/generic/os_name", http.HandlerFunc(informant.OsName))
}


