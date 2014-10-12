package osx

import (
	"net/http"
	"fmt"

	"github.com/bmizerany/pat"

	"github.com/trevrosen/target-snitch/platform/generic"
)


type OSXInformant struct{
	generic.GenericInformant
}


func New(pat *pat.PatternServeMux) (*OSXInformant){
	informant     := new(OSXInformant)
	informant.Pat = pat
	return informant
}

func (informant *OSXInformant) RegisterRoutes(){
	informant.Pat.Get("/sysctl/machdep/cpu/core_count", http.HandlerFunc(informant.sysctlMachdepCpuCoreCount))

}

// sysctlMachdepCpuCoreCount returns result of:
// $> sysctl -n machdep.cpu.core_count
func (informant *OSXInformant) sysctlMachdepCpuCoreCount(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Something coole")
}
