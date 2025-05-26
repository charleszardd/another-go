package greetings

import(
	"testing"
	"regexp"

)

func TestHelloName(t *testing.T) {
	name := "Charles"
	want := regexp.MustCompile(`/b`+name+`/b`)
	msg, _ ,_ error := Greet("Charles", 12)
	  if !want.MatchString(msg) || error != nil {
        t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, error, want)
    }
}

func TestHelloEmpty(t *testing.T) {
    msg,_,_ error := Hello("", 12)
    if msg != "" || err == nil {
        t.Errorf(`Hello("") = %q, %v, want "", error`, msg, error)
    }
}