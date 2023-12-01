package core

import (
	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
	"github.com/pidgy/obs/uptr"
)

// AddDataPath wraps EXPORT char *obs_find_data_file(const char *file).
// https://github.com/obsproject/obs-studio/blob/master/libobs/obs.h#L324
func AddDataPath(path string) error {
	return dll.OBS("obs_add_data_path", uptr.FromString(path))

}

// Initialized wraps bool obs_initialized(void).
func Initialized() (bool, error) {
	return dll.OBSbool("obs_initialized")
}

// Locale wraps obs_get_locale.
func Locale() (locale.Type, error) {
	r, err := dll.OBSuintptr("obs_get_locale")
	return locale.Type(uptr.String(r)), err
}

// SetLocale wraps void obs_set_locale(const char *locale).
func SetLocale(l locale.Type) error {
	return dll.OBS("obs_set_locale", uptr.FromString(l.String()))
}

// Shutdown wraps obs_shutdown.
func Shutdown() error {
	defer dll.Cleanup()
	return dll.OBS("obs_shutdown")
}

// Startup wraps obs_startup, use profiler.None as NULL value.
func Startup(locale locale.Type, moduleConfigPath string, ns profiler.NameStore) error {
	_, _, err := dll.Core("obs.dll")
	if err != nil {
		return err
	}

	_, _, err = dll.Core("obs-frontend-api.dll")
	if err != nil {
		return err
	}

	return dll.OBS("obs_startup", uptr.FromString(locale.String()), uptr.FromString(moduleConfigPath), uintptr(ns))
}

// Version wraps obs_get_version_string and obs_get_version.
func Version() (s string, u uint32, err error) {
	s, err = dll.OBSstring("obs_get_version_string")
	if err != nil {
		return
	}
	u, err = dll.OBSuint32("obs_get_version")
	return
}
