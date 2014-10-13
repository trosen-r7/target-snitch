# Target Snitch

(work-in-progress)

A lab system needs to be visible from the outside. TargetSnitch
provides information from the following sources and makes it available as JSON served over HTTP:

  * /proc/* (Linux)
  * sysctl (OSX)
  * tasklist.exe, WMI, etc (Windows)
  
  
## Rules/Principles

TargetSnitch must:

1.  Be able to cross compile for any target platform on a single platform - so adding Cgo stuff is verbotten. This shouldn't matter because of (2)
2. Have *zero* dependencies on the target side - this means no custom task management stuff or anything like that for getting information about system state.

## Quick Start

(on OSX)

* Install Go if you don't have it already, including setting up a
  $GOPATH variable.

* Clone this repo into $GOPATH/src/github.com/trevrosen/target-snitch
  per Go standard practice.

* Go into the dir and `go run main.go`

* That window will be bound w/ no output

* Point your browswer to http://localhost:12345

* Physical core count: http://localhost:12345/sysctl/machdep/cpu/core_count

* OS family name: http://localhost:12345/generic/os_name

* OS arch name: http://localhost:12345/generic/os_arch

