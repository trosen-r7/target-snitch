package oslinux

import(
	"testing"

	"github.com/stretchr/testify/assert"
)

var procfsParse = []struct {
	in string
	out map[string] string
}{
	{
	in:`
Name:	sshd
State:	S (sleeping)
Tgid:	860
`,
	out: map[string]string{
		"Name":"sshd",
		"State":"S (sleeping)",
		"Tgid":"860",
		},
	},
}


func TestSomeString(t *testing.T) {
	actual := SomeString()
	assert.Equal(t, "This is a thing", actual)
}


func TestParsedPairs(t *testing.T) {
	for _, test := range procfsParse{
		actual, _ := ParsedPairs([]byte(test.in), ":")
		assert.Equal(t, test.out, actual)
	}
}

