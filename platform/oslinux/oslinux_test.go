package oslinux

import(
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for key-value data separated by colons,
// similar to the output of `cat /proc/<PID>/status`
var procfsStatus = []struct {
	in string
	out []map[string] string
}{
	{
	in:`
Name:	sshd
State:	S (sleeping)
Tgid:	860
`,
	out: []map[string] string{
			{ "Name":"sshd" },
			{ "State":"S (sleeping)" },
			{ "Tgid":"860" },
		},
	},
}


// Test that ParsedPairs can parse data similar to what comes from
// `cat /proc/<PID>/status`
func TestParsedPairsWithColon(t *testing.T) {
	for _, test := range procfsStatus{
		actual, _ := ParsedPairs([]byte(test.in), ":")
		assert.Equal(t, test.out, actual)
	}
}

