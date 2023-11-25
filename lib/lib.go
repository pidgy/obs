package lib

import "syscall"

// OBS wraps obs.dll access in a single package.
var OBS = syscall.NewLazyDLL(`obs.dll`)
