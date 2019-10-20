package tagger

import (
	"fmt"
	"strings"
)

// Format ...
func Format(s string, subst map[string]interface{}) string {
	tokens := findTokens(s)

	for vv, ff := range tokens {
		ss := strings.ReplaceAll(s, `${`+vv+`:`+ff+`}`, `%`+ff)
		s = fmt.Sprintf(ss, subst[vv])
	}

	return s
}

func findTokens(s string) map[string]string {
	m := make(map[string]string)

	openers := strings.Split(s, "${")
	if len(openers) < 2 {
		return m
	}

	for _, o := range openers[1:] {
		closer := strings.Split(o, "}")
		if len(closer) < 1 {
			continue
		}

		pair := strings.Split(closer[0], ":")
		if len(pair) != 2 {
			continue
		}

		m[pair[0]] = pair[1]
	}

	return m
}
