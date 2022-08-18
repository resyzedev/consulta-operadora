package main

import (
	"consulta-operadora/funcs"
	"flag"
	"fmt"
	"html"
	"os"
	"strings"
)

var logo = `
_  _ _  _ _  _ ___  ____ ____ ____ ____ _  _ 
|  | |\ | |_/  |__] |___ |__/ [__  |  | |\ | 
|__| | \| | \_ |    |___ |  \ ___] |__| | \| 
`

func init() {
	funcs.Message(logo, "purple")
}

func main() {
	var version bool

	flag.BoolVar(&version, "version", false, "Software version.")
	flag.BoolVar(&version, "v", false, "Software version.")
	phone := flag.String("phone", "", "Phone to search.")
	flag.Parse()

	if version {
		funcs.Message("[!] version 1.0", "white")
		os.Exit(0)
	}

	if *phone == "" {
		funcs.Message("[!] no phone number found", "red")
		os.Exit(0)
	}

	if !funcs.IsValidPhone(*phone) {
		funcs.Message("[!] enter a valid phone number", "red")
		os.Exit(0)
	}

	// do post request
	data, time, status := funcs.DoPost(*phone)

	// remove html entities
	data = html.UnescapeString(data)

	// check status code
	if status != 200 {
		funcs.Message("[!] statusCode is not 200", "red")
		os.Exit(0)
	}
	fmt.Println(fmt.Sprintf("req delay => %s", time))

	if strings.Contains(data, "Aguarde") {
		time, _ := funcs.GetStr(data, `<span class="lead laranja" id="span_tempo">`, `</span>`)
		funcs.Message("[!] wait "+time+"s and try again", "yellow")
		os.Exit(0)
	}

	operadora, _ := funcs.GetStr(data, `<span class="azul lead">Operadora:</span><span class="lead laranja">`, `</span>`)
	portado, _ := funcs.GetStr(data, `<span class="azul lead">Portado:</span><span class="lead laranja">`, `</span>`)

	// remove spaces from the beginning
	operadora = strings.TrimSpace(operadora)
	portado = strings.TrimSpace(portado)

	// to lower case
	operadora = strings.ToLower(operadora)
	portado = strings.ToLower(portado)

	funcs.Message("["+*phone+"] operadora: "+operadora+" - portado: "+portado, "green")
}
