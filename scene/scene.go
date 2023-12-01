package scene

import (
	"syscall"

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
	r, err := dll.OBSuintptr("obs_scene_create", uptr.FromString(name))
	return Type(r), err
}

// Current wraps obs_source_t *obs_frontend_get_current_scene(void).
func Current() (Type, error) {
	r, err := dll.OBSuintptr("obs_frontend_get_current_scene")
	return Type(r), err
}

// IsNull returns true or false as to whether or not Item has been initialized.
func (i Item) IsNull() bool {
	return i == Item(Null)
}

// SetLocked wraps bool obs_sceneitem_set_locked(obs_sceneitem_t *item, bool locked).
func (i Item) SetLocked(b bool) (bool, error) {
	return dll.OBSbool("obs_sceneitem_set_locked", uintptr(i), uptr.FromBool(b))
}

// Locked wraps bool obs_sceneitem_locked(const obs_sceneitem_t *item).
func (i Item) Locked() (bool, error) {
	return dll.OBSbool("obs_sceneitem_locked", uintptr(i))
}

// Add wraps obs_sceneitem_t *obs_scene_add(obs_scene_t *scene, obs_source_t *source).
func (t Type) Add(s source.Type) (Item, error) {
	r, err := dll.OBSuintptr("obs_scene_add", uintptr(t), uintptr(s))
	return Item(r), err
}

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}

// Items wraps void obs_scene_enum_items(obs_scene_t *scene, bool (*callback)(obs_scene_t*, obs_sceneitem_t*, void*), void *param).
func (t Type) Items() (items []Item, err error) {
	return items, dll.OBS("obs_scene_enum_items", uintptr(t), syscall.NewCallback(
		func(scene, item, void uintptr) int {
			items = append(items, Item(item))
			return 1
		},
	))
}
