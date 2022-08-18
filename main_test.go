/* ==> go test */

package main

import (
	"consulta-operadora/funcs"
	"fmt"
	"html"
	"os"
	"strings"
	"testing"
)

// GetStr Returns empty string if no start string found
func GetStr(str string, start string, end string) (result string, found bool) {
	s := strings.Index(str, start)
	if s == -1 {
		return result, false
	}
	newS := str[s+len(start):]
	e := strings.Index(newS, end)
	if e == -1 {
		return result, false
	}
	result = newS[:e]
	return result, true
}

func TestGetStr(t *testing.T) {
	body := `<span class="lead laranja" id="span_tempo">11</span>`
	result, found := GetStr(body, `<span class="lead laranja" id="span_tempo">`, `</span>`)
	if !found {
		t.Error("string not found")
		os.Exit(0)
	}
	t.Error(result)
}

func TestContains(t *testing.T) {
	str := `<span class="lead laranja">Aguarde</span>`
	if strings.Contains(str, "Aguarde") {
		t.Error("found")
	} else {
		t.Error("not found")
	}
}

func TestChar(t *testing.T) {
	operadora := ` Telef�nica Brasil (M&oacute;vel/SMP)`
	portado := ` N&Atilde;O`
	remove_spaces_operadora := strings.TrimSpace(operadora)
	remove_html_entities_operadora := html.UnescapeString(remove_spaces_operadora)

	fmt.Println(remove_html_entities_operadora)
	fmt.Println("----")
	fmt.Println(strings.TrimSpace(html.UnescapeString(portado)))
}

func TestRegex(t *testing.T) {
	test := funcs.IsValidPhone("11975738658")
	fmt.Println(test)
}
