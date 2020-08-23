package problem0071

import (
	"strings"
)

func simplifyPath(path string) string {
	split := strings.Split(path, "/")
	var p []string
	for _, s := range split {
		switch s {
		case "":
		case ".":
		case "..":
			if len(p) > 0 {
				p = p[:len(p)-1]
			}
		default:
			p = append(p, s)
		}
	}
	return "/" + strings.Join(p, "/")
}
