package android

import (
	"github.com/bmizerany/pat"

	"github.com/trevrosen/target-snitch/informant/generic"
)

type AndroidInformant struct {
	generic.GenericInformant
}

func New(pat *pat.PatternServeMux) *AndroidInformant {
	informant := new(AndroidInformant)
	informant.Pat = pat
	return informant
}

func (informant *AndroidInformant) RegisterRoutes() {
	// something Android-y
}
