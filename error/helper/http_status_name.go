package error_helper

import (
	"net/http"
	"strings"
)

func GetHttpStatusName(status int) string {
	if httpname := http.StatusText(status); httpname != "" {
		s := strings.ToUpper(httpname)
		s = strings.ReplaceAll(s, " ", "_")
		return s
	}

	return http.StatusText(http.StatusInternalServerError)
}
