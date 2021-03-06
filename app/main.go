package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

var table = []string{
	"?", "?", "?", "?", "?", "?", "?", "?", "?", "?", //0..9
	"?", "?", "?", "?", "?", "?", "?", "?", "?", "?", //10..19
	"?", "?", "?", "?", "?", "?", "?", "?", "?", "?", //20..29
	"?", "?", "?", "?", "?", "?", "?", "?", "?", "?", //30..39
	"-.--.-", "-.--.-", "?", ".-.-.", "..-..", "-....-", ".-.-.-", "-..-.", "-----", ".----", //40..49
	"..---", "...--", "....-", ".....", "-....", "--...", "---..", "----.", "---...", "-.-.-", //50..59
	"?", "-...-", "?", "..--..", "?", ".-", "-...", "-.-.", "-..", ".", //60..69
	"..-.", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", //70..79
	"---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", //80..89
	"-.--", "--..", "?", "?", "?", "?", "?", ".-", "-...", "-.-.", //90..99
	"-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", //100..109
	"-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", //110..119
	"-..-", "-.--", "--..", "?", "?", "?", "?", "?", "?", "?", //120..129
}

func morze(c web.C, w http.ResponseWriter, _ *http.Request) {
	term := c.URLParams["term"]
	code := translate(term)
	_, _ = fmt.Fprint(w, code)
}

func gopher(c web.C, w http.ResponseWriter, _ *http.Request) {
	content, err := ioutil.ReadFile("art.txt")
	if err != nil {
		panic(err)
	}

	// Convert []byte to string and print to screen
	art := string(content)
	_, _ = fmt.Fprint(w, art)
}

func translate(term string) string {
	var result strings.Builder
	for _, c := range term {
		if c <= 122 {
			result.WriteString(table[c])
		} else {
			result.WriteString("?")
		}
		result.WriteString(" ")
	}

	if result.Len() == 0 {
		return ""
	}

	r := result.String()
	return r[:len(r)-1]
}

func main() {
	goji.Get("/gopher", gopher)
	goji.Get("/:term", morze)
	goji.Serve()
}
