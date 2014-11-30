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

### Building for platforms:

**Note: you need cross compilation support in your local Go install**

For Linux (64-bit) targets:

`GOARCH=amd64 GOOS=linux go build`

For Linux (ARM7) targets:

`GOARCH=arm GOARM=7 GOOS=linux go build`

For OSX (64-bit) targets:

`GOARCH=amd64 GOOS=darwin go build`


## Informants
Each supported operating system/platform is genericized into "Informant" types which declare URL routes and their handler functions. Each of these relative URLs will return JSON-ified responses from running the commands on the target host.

#### Generic
* Process listing: `/generic/ps` => `ps -Af`
* Target OS Name: `/generic/os_name`
* Target OS Arch: `/generic/os_arch`


#### Linux
* `/etc/issue` => `cat /etc/issue`
* `/proc/cpuinfo` => `cat /proc/cpuinfo`
* `/proc/:pid/status` => `cat /proc/:pid/status`

#### OS X
* `/sysctl/machdep/cpu/core_count` => `sysctl -n machdep.cpu.core_count`
