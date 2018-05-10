package command

import "strings"

// Flag describes a flag. Name is a short name of a flag.
// LongName is a long name of a flag. Description is
// a message for help. Value is a default value.
type Flag struct {
	Name        string
	LongName    string
	Description string
	Value       string
}

// FindFlag find a flag from a flag string.
func (fs FlagSlice) FindFlag(name string) *Flag {
	if strings.HasPrefix(name, "--") {
		for _, f := range fs {
			if f.LongName == name {
				return &f
			}
		}
	} else {
		for _, f := range fs {
			if f.Name == name {
				return &f
			}
		}
	}

	return nil
}
