package tfmt

import (
	"fmt"
	"io"
	"strings"
)

func Tprintf(format string, params map[string]interface{}) (n int, err error) {
	for key, value := range params {
		format = strings.Replace(format, "%{"+key+"}s", fmt.Sprintf("%s", value), -1)
	}
	return fmt.Printf(format)
}

func Tfprintf(w io.Writer, format string, params map[string]interface{}) (n int, err error) {
	for key, value := range params {
		format = strings.Replace(format, "%{"+key+"}s", fmt.Sprintf("%s", value), -1)
	}
	return fmt.Fprintf(w, format)
}

func Tsprintf(format string, params map[string]interface{}) string {
	for key, value := range params {
		format = strings.Replace(format, "%{"+key+"}s", fmt.Sprintf("%s", value), -1)
	}
	return fmt.Sprintf(format)
}
