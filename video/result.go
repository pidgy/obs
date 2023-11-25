package video

import "fmt"

// Result wraps integer return types.
type Result int32

const (
	Success Result = -iota
	Fail
	NotSupported
	InvalidParam
	CurrentlyActive
	ModuleNotFound
)

// String returns the string translation of a Result.
func (r Result) String() string {
	switch r {
	case Success:
		return "success"
	case Fail:
		return "fail"
	case NotSupported:
		return "not supported"
	case InvalidParam:
		return "invalid param"
	case CurrentlyActive:
		return "currently active"
	case ModuleNotFound:
		return "module not found"
	default:
		return fmt.Sprintf("unknown result %d", r)
	}
}
