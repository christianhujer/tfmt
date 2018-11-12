package tfmt

import (
	"fmt"
	"github.com/christianhujer/interceptor"
	"strings"
	"testing"
)

func AssertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Error(fmt.Sprintf("Expected: %s\nActual: %s\n", expected, actual))
	}
}

func TestTsprintf(t *testing.T) {
	expected := "foo bar"
	actual := Tsprintf("%{p1}s %{p2}s", map[string]interface{}{"p1": "foo", "p2": "bar"})
	AssertEquals(t, expected, actual)
}

func TestTsprintfDoesNotFormatAgain(t *testing.T) {
	expected := "%s bar"
	actual := Tsprintf("%{p1}s %{p2}s", map[string]interface{}{"p1": "%s", "p2": "bar"})
	AssertEquals(t, expected, actual)
}

func TestTfprintf(t *testing.T) {
	expected := "foo bar"
	builder := strings.Builder{}
	Tfprintf(&builder, "%{p1}s %{p2}s", map[string]interface{}{"p1": "foo", "p2": "bar"})
	AssertEquals(t, expected, builder.String())
}

func TestTfprintfDoesNotFormatAgain(t *testing.T) {
	expected := "%s bar"
	builder := strings.Builder{}
	Tfprintf(&builder, "%{p1}s %{p2}s", map[string]interface{}{"p1": "%s", "p2": "bar"})
	AssertEquals(t, expected, builder.String())
}

func TestTprintf(t *testing.T) {
	expected := "foo bar"
	stdout, stderr, err := interceptor.InterceptStrings(func() { Tprintf("%{p1}s %{p2}s", map[string]interface{}{"p1": "foo", "p2": "bar"}) })
	if err != nil {
		t.Error(err)
	}
	AssertEquals(t, expected, *stdout)
	AssertEquals(t, "", *stderr)
}

func TestTprintfDoesNotFormatAgain(t *testing.T) {
	expected := "%s bar"
	stdout, stderr, err := interceptor.InterceptStrings(func() { Tprintf("%{p1}s %{p2}s", map[string]interface{}{"p1": "%s", "p2": "bar"}) })
	if err != nil {
		t.Error(err)
	}
	AssertEquals(t, expected, *stdout)
	AssertEquals(t, "", *stderr)
}

//func TestTsprintfDoesNotReplaceRecursively(t *testing.T) {
//	expected := "%{p2}s bar"
//	actual := Tsprintf("%{p1}s %{p2}s", map[string]interface{}{"p1": "%{p2}s", "p2": "bar"})
//	AssertEquals(t, expected, actual)
//}
