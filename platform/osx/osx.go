package osx

import (
	"net/http"

	"github.com/bmizerany/pat"

	"github.com/trevrosen/target-snitch/platform/generic"
)


type OSXInformant struct{
	generic.GenericInformation
	generic.GenericInformant
}


func New(pat *pat.PatternServeMux) (*OSXInformant){
	informant := new(OSXInformant)
	informant.Pat = pat
	return informant
}


func (OSXInformant) OsName (w http.ResponseWriter, req *http.Request){}
