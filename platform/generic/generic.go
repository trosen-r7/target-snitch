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
	"strings"

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

// OsName returns the value of runtime.GOOS
func (informant *GenericInformant) OsName(w http.ResponseWriter, req *http.Request){
	response := JsonMarshalSingleValue(runtime.GOOS)
	io.WriteString(w, string(response))
}

// RegisterRoutes registers URL route patterns and handler functions
// for all information handed back by the Informant.
func (informant *GenericInformant) RegisterRoutes(){
	informant.Pat.Get("/", http.HandlerFunc(informant.Root))
	informant.Pat.Get("/generic/os_name", http.HandlerFunc(informant.OsName))
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

func JsonMarshalSingleValue(cmdResult interface{}) []byte {
	response := NewJsonResponse()

	switch t := cmdResult.(type){
	case bytes.Buffer:
		response.Result["value"] = strings.TrimSpace(cmdResult.String())
	case string:
		response.Result["value"] = strings.TrimSpace(cmdResult)
	}

	jsonBytes, err := json.Marshal(response.Result)
	if err != nil {
		log.Fatal(err)
	}
	return jsonBytes
}

func RunCommand(cmd *exec.Cmd) (bytes.Buffer){
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil{
		log.Fatal(err)
	}
	return out
}
