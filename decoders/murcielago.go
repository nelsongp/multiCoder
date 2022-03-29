package decoders

import "fmt"

type Murcielago struct{}

func (m *Murcielago) Notify(name string) {
	codeMurcielago(name)
}

func codeMurcielago(name string) {
	fmt.Println(name)
}
