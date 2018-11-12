package tfmt

import (
	"fmt"
	"io"
	"strings"
)

func formatString(format string, params map[string]interface{}) string {
	for key, value := range params {
		format = strings.Replace(format, "%{"+key+"}s", fmt.Sprintf("%s", value), -1)
	}
	return format
}

func Tprintf(format string, params map[string]interface{}) (n int, err error) {
	return fmt.Print(formatString(format, params))
}

func Tfprintf(w io.Writer, format string, params map[string]interface{}) (n int, err error) {
	return fmt.Fprint(w, formatString(format, params))
}

func Tsprintf(format string, params map[string]interface{}) string {
	return fmt.Sprint(formatString(format, params))
}
