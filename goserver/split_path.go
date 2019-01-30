package goserver

import (
	"path"
	"strings"
)

// SplitPath splits off the first component of url, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
func SplitPath(url string) (head, tail string) {
	url = path.Clean("/" + url)
	i := strings.Index(url[1:], "/") + 1
	if i <= 0 {
		return url[1:], "/"
	}
	return url[1:i], url[i:]
}
