package httpHelper

import "strings"

func RetriveIdFromUrlPath(path string) string {
	return strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
}
