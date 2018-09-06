package str

import (
	"strings"
)

// GetFirstLastName returns the first and last name from a string,
// removing extra spaces.
func GetFirstLastName(s string) (string, string) {
	names := strings.Fields(s)
	if len(names) == 0 {
		return "", ""
	}
	if len(names) == 1 {
		return names[0], ""
	}

	return names[0], strings.Join(names[1:], " ")
}
