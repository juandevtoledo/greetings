package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, error := Hello("Juan")
	if !want.MatchString(msg) || error != nil {
		t.Fatalf(`Hello("juan") = %q,%v, quiere coincidencia para %#q, nil`, msg, error, want)
	}
}

func TestHelloEmpty(t *testing.T) {

	msg, error := Hello("")
	if msg != "" || error == nil {
		t.Fatalf(`Hello("") = %q,%v, quiere "", error`, msg, error)
	}
}
