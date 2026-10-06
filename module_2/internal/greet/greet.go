package greet

import "strings"

func Hello(name string ) string {
	clean := normalizeName(name)

	return "Hello ," + clean
}

func normalizeName(name string) string {
	name = strings.TrimSpace(name)

	if name == "" {
		return "Guest"
	}

	return strings.ToUpper(name)
}