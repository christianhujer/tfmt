package tfmt

import (
	"fmt"
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

//func TestTsprintfDoesNotReplaceRecursively(t *testing.T) {
//	expected := "%{p2}s bar"
//	actual := Tsprintf("%{p1}s %{p2}s", map[string]interface{}{"p1": "%{p2}s", "p2": "bar"})
//	AssertEquals(t, expected, actual)
//}
