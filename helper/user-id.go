package helper

import "strings"

// MapUserID changes the id to desired pattern
func MapUserID(id *string) {
	*id = strings.ToUpper(*id)
}
