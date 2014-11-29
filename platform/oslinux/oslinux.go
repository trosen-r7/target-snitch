package oslinux

import(
	"net/http"
	"io"
	"bytes"
	"encoding/json"
	"os/exec"
	"log"
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

// procCpuInfo provides information from `cat /proc/cpuinfo`
func (informant *LinuxInformant)procCpuInfo(w http.ResponseWriter, req *http.Request)  {
	cmd       := exec.Command("cat", "/proc/cpuinfo")
	cmdResult := generic.RunCommand(cmd)

	parsedLines, err := ParsedPairs(cmdResult.Bytes(), ":")
	if err != nil {
		log.Fatalln(err)
	}

	responseJson, err := json.Marshal(parsedLines)
	if err != nil {
		log.Fatalln(err)
	}

	io.WriteString(w, string(responseJson))
}

// procPidStatus provides information from `cat /proc/:pid/status`
func (informant *LinuxInformant)procPidStatus(w http.ResponseWriter, req *http.Request) {
	argString := "/proc/" + req.URL.Query().Get(":pid") + "/status"
	cmd       := exec.Command("cat", argString)
	cmdResult := generic.RunCommand(cmd)

	parsedLines, err := ParsedPairs(cmdResult.Bytes(), ":")
	if err != nil {
		log.Fatalln(err)
	}

	responseJson, err := json.Marshal(parsedLines)
	if err != nil {
		log.Fatalln(err)
	}

	io.WriteString(w, string(responseJson))
}


// RegisterRoutes registers URL route patterns and handler functions
// for all information handed back by the Informant.
func (informant *LinuxInformant) RegisterRoutes(){
	informant.Pat.Get("/proc/cpuinfo", http.HandlerFunc(informant.procCpuInfo))
	informant.Pat.Get("/proc/:pid/status", http.HandlerFunc(informant.procPidStatus))
}

// ParsedPairs splits a byte slice into lines and each line on a separator, returning
// a slice of maps corresponding to the splits, and an error if the lines contain 
// more than 2 items after split.
func ParsedPairs(cmdOriginalOutput []byte, separator string) (outputArray []map[string] string, err error) {
	parsedLines := bytes.Split(cmdOriginalOutput, []byte("\n"))

	for _, line := range parsedLines{
		if len(line) == 0 { continue }

		pair := bytes.Split(line, []byte(separator))
		if len(pair) != 2 {
			// return error and empty map
		}
		pairMap      := make(map[string]string)
		key          := string(bytes.TrimSpace(pair[0]))
		value        := string(bytes.TrimSpace(pair[1]))
		pairMap[key] = value
		outputArray  = append(outputArray, pairMap)
	}

	return outputArray, err
}

