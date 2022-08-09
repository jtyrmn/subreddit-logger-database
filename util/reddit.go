package util

import "regexp"

// the listings in the database must store IDs in the form of t-_------, or
// regex ^t[1-6]_[a-z0-9]{6}$
func IsValidID(ID string) bool {
	result, _ := regexp.MatchString("^t[1-6]_[a-z0-9]{6}$", string(ID))
	return result
}
