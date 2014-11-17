// Used to return things that can be had from the runtime package
// or otherwise calculated in a generic way across supported OSes.
package generic

import(
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	_"fmt"
	"strings"

	"github.com/bmizerany/pat"
)


// The fields in a 'ps -f' shell command
var psfFields = [...]string {"UID", "PID", "PPID", "C", "STIME", "TTY", "TIME", "CMD", "ARGV"}


// GenericInformation implements functions that return OS-agnostic stats
type GenericInformation interface {
	RegisterRoutes()
	// OsName returns the value of runtime.GOOS
	//OsName(w http.ResponseWriter, req *http.Request)
	// OsArch returns the value of runtime.GOARCH
	//OsArch()

}

// GenericInformant provides a minimal interface for dealing w/ the wire
// All platform Informant types will embed this type.
type GenericInformant struct {
	Pat *pat.PatternServeMux
}

// JsonResponse holds a generic structure for marshalling JSON
type JsonResponse struct {
	Result map[string]interface{}
}

// NewJsonResponse initializes a JsonResponse
// Use this when you need to manually set up a bunch of arbitrary JSON
// in a function and need a return value.
func NewJsonResponse() *JsonResponse {
	return &JsonResponse{Result: make(map[string]interface{})}
}

// New creates a new GenericInformant
func New(pat *pat.PatternServeMux) (*GenericInformant){
	informant     := new(GenericInformant)
	informant.Pat = pat
	return informant
}

// OsArch returns the value of runtime.GOARCH
func (informant *GenericInformant) OsArch(w http.ResponseWriter, req *http.Request){
	response := JsonMarshalSingleValue(runtime.GOARCH)
	io.WriteString(w, string(response))
}

// OsName returns the value of runtime.GOOS
func (informant *GenericInformant) OsName(w http.ResponseWriter, req *http.Request){
	response := JsonMarshalSingleValue(runtime.GOOS)
	io.WriteString(w, string(response))
}

// Ps returns the output of (ps -Af), returning JSON with keys
// UID, PID, PPID, C, STIME, TTY, TIME, CMD and ARGV
// all but ARGV are listed as the headers when running `ps -f`. ARGV is custom
// to target-snitch in order to represent the ARGV passed to the CMD.
func (informant *GenericInformant) Ps(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("ps", "-Af")

	cmdResult     := RunCommand(cmd)
	responseLines := bytes.Split(cmdResult.Bytes(), []byte("\n"))
	responseArray := []map[string]string{}

	for i, psLine := range responseLines {
		if i == 0 { continue } // skip header line
		var lineMap  = make(map[string]string)
		for j, fieldBytes := range bytes.Fields(psLine) {
			if j < len(psfFields){
				lineMap[psfFields[j]] = string(fieldBytes)
			}
		}
		if len(lineMap) > 0 {
			responseArray = append(responseArray, lineMap)
		}
	}

	responseJSON, error     := json.Marshal(responseArray)
	if error != nil {
		log.Fatalln(error)
	}
	io.WriteString(w, string(responseJSON))
}


// RegisterRoutes registers URL route patterns and handler functions
// for all information handed back by the Informant.
func (informant *GenericInformant) RegisterRoutes(){
	informant.Pat.Get("/", http.HandlerFunc(informant.Root))

	informant.Pat.Get("/generic/os_arch", http.HandlerFunc(informant.OsArch))
	informant.Pat.Get("/generic/os_name", http.HandlerFunc(informant.OsName))
	informant.Pat.Get("/generic/ps", http.HandlerFunc(informant.Ps))
}

func (informant *GenericInformant) Root(w http.ResponseWriter, req *http.Request){
	responseString := `
 _______  _______  ______    _______  _______  _______    _______  __    _  ___   _______  _______  __   __
|       ||   _   ||    _ |  |       ||       ||       |  |       ||  |  | ||   | |       ||       ||  | |  |
|_     _||  |_|  ||   | ||  |    ___||    ___||_     _|  |  _____||   |_| ||   | |_     _||       ||  |_|  |
  |   |  |       ||   |_||_ |   | __ |   |___   |   |    | |_____ |       ||   |   |   |  |       ||       |
  |   |  |       ||    __  ||   ||  ||    ___|  |   |    |_____  ||  _    ||   |   |   |  |      _||       |
  |   |  |   _   ||   |  | ||   |_| ||   |___   |   |     _____| || | |   ||   |   |   |  |     |_ |   _   |
  |___|  |__| |__||___|  |_||_______||_______|  |___|    |_______||_|  |__||___|   |___|  |_______||__| |__|


	     (=) Consult docs for available Informants (=)
	`
	io.WriteString(w, responseString)
}

// JsonMarshallSingleValue takes either a bytes.Buffer or a string type
// and places it into a JSON-encoded string at the key "value".
func JsonMarshalSingleValue(cmdResult interface{}) []byte {
	response := NewJsonResponse()

	switch t := cmdResult.(type){
	case bytes.Buffer:
		response.Result["value"] = strings.TrimSpace(t.String())
	case string:
		response.Result["value"] = strings.TrimSpace(t)
	}

	jsonBytes, err := json.Marshal(response.Result)
	if err != nil {
		log.Fatal(err)
	}
	return jsonBytes
}

// RunCommand uses the os/exec package to run a shell
// command on the system.
func RunCommand(cmd *exec.Cmd) (bytes.Buffer){
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil{
		log.Fatal(err)
	}
	return out
}

