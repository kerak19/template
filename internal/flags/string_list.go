package flags

import "strings"

// StringList is an flag type allowing to pass multiple strings as one argument
type StringList []string

func (sl StringList) String() string {
	return strings.Join(sl, ",")
}

// Set set's provided string s as a list of strings
func (sl *StringList) Set(s string) error {
	*sl = strings.Split(s, ",")
	return nil
}
