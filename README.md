# Target Snitch

(work-in-progress)

A lab system needs to be visible from the outside. TargetSnitch will
provide information from:

  * /proc/* (Linux)
  * sysctl (OSX)
  * tasklist.exe, WMI, etc (Windows)
  
  
## Rules/Principles

TargetSnitch must:

1.  Be able to cross compile for any target platform on a single platform - so adding Cgo stuff is verbotten. This shouldn't matter because of (2)
2. Have *zero* dependencies on the target side - this means no custom task management stuff or anything like that for getting information about system state.
