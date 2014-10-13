package osx

import (
	"net/http"
	"os/exec"
	"io"

	"github.com/bmizerany/pat"

	"github.com/trevrosen/target-snitch/platform/generic"
)


type OSXInformant struct{
	generic.GenericInformant
}


// New creates a new OSXInformant
func New(pat *pat.PatternServeMux) (*OSXInformant){
	informant     := new(OSXInformant)
	informant.Pat = pat
	return informant
}

// RegisterRoutes registers URL route patterns and handler functions
// for all information handed back by the Informant.
func (informant *OSXInformant) RegisterRoutes(){
	informant.Pat.Get("/sysctl/machdep/cpu/core_count", http.HandlerFunc(informant.sysctlMachdepCpuCoreCount))

}

// sysctlMachdepCpuCoreCount returns result of:
// $> sysctl -n machdep.cpu.core_count
func (informant *OSXInformant) sysctlMachdepCpuCoreCount(w http.ResponseWriter, req *http.Request) {
	cmd				:= exec.Command("sysctl", "-n", "machdep.cpu.core_count")

	cmdResult := generic.RunCommand(cmd)
	io.WriteString(w, string(generic.JsonMarshalSingleValue(cmdResult)))
}



