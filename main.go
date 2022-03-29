package main

import (
	"bufio"
	"fmt"
	"github.com/nelsongp/multiCoder/facade"
	"log"
	"os"
	"strings"
)

func main() {
	m := readMessage()
	f := facade.New()
	f.Encoders(m)
}

func readMessage() string {
	fmt.Print("Digite un Nombre: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("no se pudo leer el mensaje del usuario: #{err}")
	}
	text = strings.TrimSuffix(text, "\n")
	return text
}
