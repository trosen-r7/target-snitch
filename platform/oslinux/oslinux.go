package oslinux

import(
	"net/http"
	"io"

	"github.com/bmizerany/pat"
	"github.com/trevrosen/target-snitch/platform/generic"
)

type LinuxInformant struct {
	generic.GenericInformant
}

// New creates a new Linux informant
func New(pat *pat.PatternServeMux) (*LinuxInformant) {
	informant := new(LinuxInformant)
	informant.Pat = pat
	return informant
}

func (informant *LinuxInformant)procCpuInfo(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "something")
}


// RegisterRoutes registers URL route patterns and handler functions
// for all information handed back by the Informant.
func (informant *LinuxInformant) RegisterRoutes(){
	informant.Pat.Get("/proc/cpuinfo", http.HandlerFunc(informant.procCpuInfo))

}

