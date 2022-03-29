package decoders

import "fmt"

type Binario struct{}

func (b *Binario) Notify(name string) {
	codeBinario(name)
}

func codeBinario(name string) {
	fmt.Println(name)
}
