package decoders

import (
	"fmt"
	"strings"
)

type Morse struct{}

func (m *Morse) Notify(name string) {
	codeMorse(name)
}

/*CodeMorse funcion que recibe un string y lo codifica a Morse*/
func codeMorse(name string) {
	n := strings.ToLower(name)
	var s []string
	for _, l := range n {
		s = append(s, searchLetter(string(l)))
	}
	fmt.Println(strings.Join(s[:], ""))
}

func searchLetter(letra string) string {
	m := make(map[string]string)
	m["a"] = ".-"
	m["b"] = "-..."
	m["c"] = ".-.-"
	m["d"] = "-.."
	m["e"] = "."
	m["f"] = "..-."
	m["g"] = "--."
	m["h"] = "...."
	m["i"] = ".."
	m["j"] = ".---"
	m["k"] = "-.-"
	m["l"] = "-.-"
	m["m"] = "--"
	m["n"] = "-.."
	m["ñ"] = "--.--"
	m["o"] = "---"
	m["p"] = ".--."
	m["q"] = "--.-"
	m["r"] = ".-."
	m["s"] = "..."
	m["t"] = "-"
	m["u"] = "..-"
	m["v"] = "...-"
	m["w"] = ".--"
	m["x"] = "-..-"
	m["y"] = "-.--"
	m["z"] = "--.."

	return m[letra]
}
