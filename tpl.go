package strtpl

import (
	"fmt"
	"strings"
)

type TPL struct {
	raw    string
	tokens []string
}

func NewTPL(raw string) *TPL {
	t := &TPL{raw: raw}
	t.tokens = strings.Split(raw, " ")
	return t
}

func (t *TPL) Format(args map[string]interface{}) string {
	tmp := make([]string, len(t.tokens))
	copy(tmp, t.tokens)
	var changes []struct {
		Index int
		Value string
	}
	for k, v := range args {
		for i, token := range tmp {
			if token == k {
				changes = append(changes, struct {
					Index int
					Value string
				}{i, fmt.Sprint(v)})
			}
		}
	}
	for _, change := range changes {
		tmp[change.Index] = change.Value
	}
	return strings.Join(tmp, " ")
}
