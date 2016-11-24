package strflag

import (
	"fmt"
	"strings"
)

// StringSlice is a custom type for converting comma-separated command line argument to a string slice.
type StringSlice []string

// String returns the string representation of the flag value.
func (s *StringSlice) String() string {
	return fmt.Sprint([]string(*s))
}

// Set sets a new value for the flag.
func (s *StringSlice) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}
