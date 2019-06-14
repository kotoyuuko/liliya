package config

import (
	"reflect"

	"github.com/Unknwon/goconfig"
)

// File represents a config file
type File struct {
	configFile *goconfig.ConfigFile
	section    string
	key        string
	defaultVal interface{}
}

// Load loads config file
func Load(path string) (*File, error) {
	cfg, err := goconfig.LoadConfigFile(path)
	if err != nil {
		return nil, err
	}

	f := new(File)
	f.configFile = cfg

	return f, nil
}

// Section set the section to read
func (f *File) Section(section string) *File {
	f.section = section
	return f
}

// Key set the key to read
func (f *File) Key(key string) *File {
	f.key = key
	return f
}

// Default set the default value to return
func (f *File) Default(defaultVal interface{}) *File {
	f.defaultVal = defaultVal
	return f
}

func isNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func (f *File) sectionSet() bool {
	return f.section != ""
}

func (f *File) keySet() bool {
	return f.key != ""
}

func (f *File) defaultSet() bool {
	return !isNil(f.defaultVal)
}

// String always returns string value without error
func (f *File) String() string {
	if !f.sectionSet() || !f.keySet() {
		return ""
	}

	if f.defaultSet() {
		return f.configFile.MustValue(f.section, f.key, f.defaultVal.(string))
	}

	return f.configFile.MustValue(f.section, f.key, "")
}

// Bool always returns bool value without error
func (f *File) Bool() bool {
	if !f.sectionSet() || !f.keySet() {
		return false
	}

	if f.defaultSet() {
		return f.configFile.MustBool(f.section, f.key, f.defaultVal.(bool))
	}

	return f.configFile.MustBool(f.section, f.key, false)
}

// Float64 always returns float64 value without error
func (f *File) Float64() float64 {
	if !f.sectionSet() || !f.keySet() {
		return 0.0
	}

	if f.defaultSet() {
		return f.configFile.MustFloat64(f.section, f.key, f.defaultVal.(float64))
	}

	return f.configFile.MustFloat64(f.section, f.key, 0.0)
}

// Int always returns int value without error
func (f *File) Int() int {
	if !f.sectionSet() || !f.keySet() {
		return 0
	}

	if f.defaultSet() {
		return f.configFile.MustInt(f.section, f.key, f.defaultVal.(int))
	}

	return f.configFile.MustInt(f.section, f.key, 0)
}

// Int64 always returns int64 value without error
func (f *File) Int64() int64 {
	if !f.sectionSet() || !f.keySet() {
		return 0
	}

	if f.defaultSet() {
		return f.configFile.MustInt64(f.section, f.key, f.defaultVal.(int64))
	}

	return f.configFile.MustInt64(f.section, f.key, 0)
}
