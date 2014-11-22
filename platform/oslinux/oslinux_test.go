package oslinux

import(
	"testing"
)

var procfsParse = []struct {
	in string
	out string
}{
	in:`
Name:	sshd
State:	S (sleeping)
Tgid:	860
Ngid:	0
Pid:	860
PPid:	1
TracerPid:	0
Uid:	0	0	0	0
Gid:	0	0	0	0
FDSize:	64
Groups:
VmPeak:	   61384 kB
VmSize:	   61364 kB
`,
	out: "something"}


func TestColonProcess(t *testing.T) {


}

