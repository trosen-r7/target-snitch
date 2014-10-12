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
	RegisterRoutes()
	// OsName returns the value of runtime.GOOS
	//OsName(w http.ResponseWriter, req *http.Request)
	// OsArch returns the value of runtime.GOARCH
	//OsArch()

}

// GenericInformant provides a minimal interface for dealing w/ the wire
type GenericInformant struct {
	Pat *pat.PatternServeMux
}

func New(pat *pat.PatternServeMux) (*GenericInformant){
	informant     := new(GenericInformant)
	informant.Pat = pat
	return informant
}

// OsName returns the value of runtime.GOOS
func (informant *GenericInformant) OsName(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "" + runtime.GOOS + "\n")
}

func (informant *GenericInformant) RegisterRoutes(){
	informant.Pat.Get("/generic/os_name", http.HandlerFunc(informant.OsName))
}

