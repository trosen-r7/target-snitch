package standard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test data for key-value return types from
// procfs - such as /proc/cpuinfo anf /proc/<PID>/status
var procfsValidKeyValue = []struct {
	in  string
	out map[string]string
}{

	// /proc/<PID>/status`
	{
		in: `
Name:	sshd
State:	S (sleeping)
Tgid:	860
`,
		out: map[string]string{
			"Name":  "sshd",
			"State": "S (sleeping)",
			"Tgid":  "860",
		},
	},

	// /proc/cpuinfo
	{
		in: `
processor	: 0
vendor_id	: GenuineIntel
cpu family	: 6
model		: 42
model name	: Intel(R) Core(TM) i7-2620M CPU @ 2.70GHz
power management	:
`,
		out: map[string]string{
			"processor":        "0",
			"vendor_id":        "GenuineIntel",
			"cpu family":       "6",
			"model":            "42",
			"model name":       "Intel(R) Core(TM) i7-2620M CPU @ 2.70GHz",
			"power management": "",
		},
	},
}

// Test that ParsedPairs can parse procfs key value pairs
func TestParsedPairsWithColon(t *testing.T) {
	for _, test := range procfsValidKeyValue {
		actual, _ := ParsedPairs([]byte(test.in), ":")
		assert.Equal(t, test.out, actual)
	}
}
