package scene

import (
	"syscall"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/source"
	"github.com/pidgy/obs/uptr"
)

type (
	// Type wraps obs_scene_t.
	Type uintptr

	// Item wraps obs_sceneitem_t.
	Item uintptr
)

const (
	// Null represents a nil obs_scene_t.
	Null = Type(0)
)

// New wraps obs_scene_t *obs_scene_create(const char *name).
func New(name string) (Type, error) {
	r, _, err := dll.OBS.NewProc("obs_scene_create").Call(
		uptr.FromString(name),
	)
	if err != syscall.Errno(0) {
		return Null, errors.Wrap(err, "obs_scene_create")
	}
	return Type(r), nil
}

// Current wraps obs_source_t *obs_frontend_get_current_scene(void).
func Current() (Type, error) {
	r, _, err := dll.Frontend.NewProc("obs_frontend_get_current_scene").Call()
	if err != syscall.Errno(0) {
		return Null, errors.Wrap(err, "obs_frontend_get_current_scene")
	}
	return Type(r), nil
}

// IsNull returns true or false as to whether or not Item has been initialized.
func (i Item) IsNull() bool {
	return i == Item(Null)
}

// SetLocked wraps bool obs_sceneitem_set_locked(obs_sceneitem_t *item, bool locked).
func (i Item) SetLocked(b bool) (bool, error) {
	r, _, err := dll.OBS.NewProc("obs_sceneitem_set_locked").Call(
		uintptr(i),
		uptr.FromBool(b),
	)
	if err != syscall.Errno(0) {
		return false, errors.Wrap(err, "obs_sceneitem_set_locked")
	}
	return uptr.Bool(r), nil
}

// Locked wraps bool obs_sceneitem_locked(const obs_sceneitem_t *item).
func (i Item) Locked() (bool, error) {
	r, _, err := dll.OBS.NewProc("obs_sceneitem_locked").Call(
		uintptr(i),
	)
	if err != syscall.Errno(0) {
		return false, errors.Wrap(err, "obs_sceneitem_locked")
	}
	return uptr.Bool(r), nil
}

// Add wraps obs_sceneitem_t *obs_scene_add(obs_scene_t *scene, obs_source_t *source).
func (t Type) Add(s source.Type) (Item, error) {
	r, _, err := dll.OBS.NewProc("obs_scene_add").Call(
		uintptr(t),
		uintptr(s),
	)
	if err != syscall.Errno(0) {
		return Item(Null), errors.Wrap(err, "obs_scene_add")
	}
	return Item(r), nil
}

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}

// Items wraps void obs_scene_enum_items(obs_scene_t *scene, bool (*callback)(obs_scene_t*, obs_sceneitem_t*, void*), void *param).
func (t Type) Items() ([]Item, error) {
	items := []Item{}

	callback := syscall.NewCallback(
		func(scene, item, void uintptr) int {
			items = append(items, Item(item))
			return 1
		},
	)

	_, _, err := dll.OBS.NewProc("obs_scene_enum_items").Call(
		uintptr(t),
		callback,
	)
	if err != syscall.Errno(0) {
		return nil, errors.Wrap(err, "obs_scene_enum_items")
	}

	return items, nil
}
