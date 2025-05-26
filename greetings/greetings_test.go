package greetings

import(
	"testing"
	"regexp"

)

func TestHelloName(t *testing.T) {
	name := "Charles"
	want := regexp.MustCompile(`/b`+name+`/b`)
	msg, error, _,_ := Greet("Charles", 12)
	  if !want.MatchString(msg) || error != nil {
        t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, error, want)
    }
}

func TestHelloEmpty(t *testing.T) {
    msg,error,_,_ := Greet("", 12)
    if msg != "" || error == nil {
        t.Errorf(`Hello("") = %q, %v, want "", error`, msg, error)
    }
}