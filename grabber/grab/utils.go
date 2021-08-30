package grab

import (
	"net/url"
	"strings"
)

func isValidFormat(u *url.URL) bool {
	path := u.Path
	resp := strings.HasSuffix(path, ".png") || strings.HasSuffix(path, ".jpg")
	resp = resp || strings.HasSuffix(path, ".jpeg") || strings.HasSuffix(path, ".gif")
	return resp
}
