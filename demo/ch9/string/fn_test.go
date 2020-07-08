package string

import (
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
}
