package web

import "strings"

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func MakePath(method string, parts ...string) string {
	var sb strings.Builder
	sb.WriteString(method)
	sb.WriteByte(' ')
	for _, part := range parts {
		sb.WriteByte('/')
		sb.WriteString(part)
	}
	sb.WriteByte('/')
	return sb.String()
}
