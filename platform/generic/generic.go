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
	"fmt"
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

// Ps returns the output of (ps -f), which contains
// UID, PID, PPID, C, STIME, TTY, TIME, CMD and the ARGV for CMD
func (informant *GenericInformant) Ps(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("ps", "-f")

	cmdResult     := RunCommand(cmd)
	responseLines := bytes.Split(cmdResult.Bytes(), []byte("\n"))
	var responseArray = make([]map[string]string)

	for i, psLine := range responseLines {
		if i == 0 { continue } // skip header line
		var lineMap map[string]string
		for j, fieldBytes := range bytes.Fields(psLine) {
			lineMap[psfFields[j]] = string(fieldBytes)
		}
		append(responseLines, lineMap)
	}

	fmt.Println(responseLines)

	// split on \n
	// remove first line
	// split on spaces and remove empties
	// parse remaining array into array of maps



	//response     := JsonMarshalMap(cmdResult)
	io.WriteString(w, "TESTING: check console")
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

// JsonMarshalStringMap will take any map of string keys and values and turn it into JSON
//func JsonMarshalStringMap(map[string]string) []byte {
//}


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

