package main

import (
  "fmt"
	"log"
	"github.com/bmizerany/pat"
	"net/http"
	"runtime"

	"github.com/trevrosen/target-snitch/platform/generic"
	"github.com/trevrosen/target-snitch/platform/osx"
	"github.com/trevrosen/target-snitch/platform/linux"
)



func main() {
	m := pat.New()

	// Generic is always registered, as the routes in there
	// only provide things from the Go runtime package.
	genericInformant := generic.New(m)
	genericInformant.RegisterRoutes()

	// Load OS-specific routes
	switch runtime.GOOS {
	case "darwin":
		osxInformant := osx.New(m)
		osxInformant.RegisterRoutes()
	}

  fmt.Println("[+] TargetSnitch is listening on port 12345")

	http.Handle("/", m)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}




