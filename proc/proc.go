package proc

import (
	"github.com/pidgy/obs/call"
	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/uptr"
)

type (
	// Handler wraps proc_handler_t.
	Handler uintptr
)

const (
	// Null represents a nil proc_handler_t pointer.
	Null = Handler(0)
)

// Call wraps bool proc_handler_call(proc_handler_t *handler, const char *name, calldata_t *params).
func (h Handler) Call(name string, data call.Data) (bool, error) {
	return dll.OBSCallBool("proc_handler_call", uintptr(h), uptr.FromString(name), uintptr(data))
}

// IsNull returns true or false as to whether or not Handler has been initialized.
func (h Handler) IsNull() bool {
	return h == Null
}
