package main

import (
	"flag"
	"fmt"
	"github.com/bmizerany/pat"
	"log"
	"net/http"
	"runtime"

	"github.com/trevrosen/target-snitch/platform/generic"
	"github.com/trevrosen/target-snitch/platform/oslinux"
	"github.com/trevrosen/target-snitch/platform/osx"
)

var listenerPort string

func init() {
	flag.StringVar(&listenerPort, "port", "12345", "port that TargetSnitch listens on")
	flag.Parse()
}

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
	case "linux":
		linuxInformant := oslinux.New(m)
		linuxInformant.RegisterRoutes()
	}

	fmt.Println("[+] TargetSnitch is listening on port", listenerPort)
	portString := ":" + listenerPort

	http.Handle("/", m)
	err := http.ListenAndServe(portString, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
