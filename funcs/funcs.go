package funcs

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var colors = map[string]string{
	"default": "\033[0m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"purple":  "\033[35m",
	"white":   "\033[97m",
}

func IsValidPhone(result string) bool {
	found, _ := regexp.MatchString(`^\(?\d{2}\)?[\s-]?[\s9]?\d{4}-?\d{4}$`, result)
	if found {
		return true
	} else {
		return false
	}
}

func Message(message, colorType string) {
	fmt.Println(colors[colorType] + message + colors["default"])
}

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

func DoPost(phone string) (string, time.Duration, int) {
	start := time.Now()

	formData := url.Values{}
	formData.Set("tipo", "consulta")
	formData.Set("numero", phone)

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://consultaoperadora.com.br/site2015/resposta.php", strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	duration := time.Since(start)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data), duration, resp.StatusCode
}

func DoGet() (string, time.Duration, int) {
	// get current time
	start := time.Now()

	resp, err := http.Get("........")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	duration := time.Since(start)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data), duration, resp.StatusCode
}
