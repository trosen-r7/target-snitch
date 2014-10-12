// Used to return things that can be had from the runtime package
// or otherwise calculated in a generic way across supported OSes.
package generic

import(
	"io"
	"net/http"
	"runtime"
	"github.com/bmizerany/pat"
)


// GenericInformation implements functions that return OS-agnostic stats
type GenericInformation interface {
	// OsName returns the value of runtime.GOOS
	OsName(w http.ResponseWriter, req *http.Request)
	// OsArch returns the value of runtime.GOARCH
	//OsArch()
	// CoreCount returns the number of cores on the machine
	//CoreCount()
	// ThreadCount returns the number of threads that can be run in parallel on the hardware
	//ThreadCount()
}

// GenericInformant provides a minimal interface for dealing w/ the wire
type GenericInformant struct {
	Pat *pat.PatternServeMux
}


// OsName returns the value of runtime.GOOS
func (GenericInformant) OsName(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "" + runtime.GOOS + "\n")
}
