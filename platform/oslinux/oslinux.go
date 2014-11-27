package oslinux

import(
	"net/http"
	"io"
	"bytes"
	_"fmt"

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


func ParsedPairs(cmdOriginalOutput []byte, separator string) (outputArray []map[string] string, err error) {
	parsedLines := bytes.Split(cmdOriginalOutput, []byte("\n"))

	for _, line := range parsedLines{
		if len(line) == 0 { continue }
		pair := bytes.Split(line, []byte(separator))
		if len(pair) != 2 {
			// return error and empty map
		}
		pairMap := make(map[string]string)
		pairMap[string(pair[0])] = string(bytes.TrimSpace(pair[1]))
		outputArray = append(outputArray, pairMap)
	}

	return outputArray, err
}

