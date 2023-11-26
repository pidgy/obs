package dshow

import (
	"fmt"
	"syscall"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/module"
	"github.com/pidgy/obs/uptr"
)

// Module is a module.Module that wraps the win-dshow type.
type Module struct {
	module.Module // Satisfies the plugin.Plugin contract.
}

const (
	file = "win-dshow.dll"
)

var (
	libdshow = dll.File(file)
)

// New returns a loaded win-dshow plugin.
func New() (*Module, error) {
	p := new(Module)

	_, _, err := dll.Module(file)
	if err != nil {
		return nil, err
	}

	return p, p.load()
}

// Close wraps obs_module_<free|unload>_* functions.
// Callers should unwrap error values to see which procedures were not implemented.
func (p *Module) Close() (werr error) {
	errs := []error{}

	if libdshow.NewProc("obs_module_free_locale").Find() == nil {
		_, _, err := libdshow.NewProc("obs_module_free_locale").Call()
		if err != syscall.Errno(0) {
			if werr == nil {
				werr = err
			}
			werr = errors.Wrap(werr, "obs_module_free_locale")
		}
	} else {
		werr = errors.Wrap(module.ErrNotImplemented, "obs_module_free_locale")
	}

	if libdshow.NewProc("obs_module_unload").Find() == nil {
		_, _, err := libdshow.NewProc("obs_module_unload").Call()
		if err != syscall.Errno(0) {
			if werr == nil {
				werr = err
			}
			werr = errors.Wrap(werr, "obs_module_unload")
		}
	} else {
		werr = errors.Wrap(module.ErrNotImplemented, "obs_module_unload")
	}

	if werr == nil && len(errs) > 0 {
		werr = errors.New("unimplemented procedures")
		for _, err := range errs {
			werr = errors.Wrap(werr, err.Error())
		}
	}

	return
}

// Description wraps obs_module_description.
func (p *Module) Description() (string, error) {
	r, _, err := libdshow.NewProc("obs_module_description").Call()
	if err != syscall.Errno(0) {
		return "", errors.Wrap(err, "obs_module_description")
	}

	return uptr.String(r), nil
}

// SetLocale wraps obs_module_set_locale.
func (p *Module) SetLocale(l locale.Type) error {
	_, _, err := libdshow.NewProc("obs_module_load").Call(
		uptr.FromString(l.String()),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_module_load")
	}
	return nil
}

// Version wraps obs_module_ver.
func (p *Module) Version() (string, error) {
	r, _, err := libdshow.NewProc("obs_module_ver").Call()
	if err != syscall.Errno(0) {
		return "", errors.Wrap(err, "obs_module_ver")
	}
	return uptr.Version(uint32(r)), nil
}

// obs_module_description = obs_module_description
// obs_module_free_locale = obs_module_free_locale
// obs_module_get_string = obs_module_get_string
// obs_module_load = obs_module_load
// obs_module_set_locale = obs_module_set_locale
// obs_module_set_pointer = obs_module_set_pointer
// obs_module_ver = obs_module_ver

// Description wraps obs_module_load.
func (p *Module) load() error {
	r, _, err := libdshow.NewProc("obs_module_load").Call()
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_module_load")
	}

	if !uptr.Bool(r) {
		return fmt.Errorf("failed to load module")
	}

	return nil
}
