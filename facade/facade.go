package facade

import (
	"github.com/nelsongp/multiCoder/decoders"
	"github.com/nelsongp/multiCoder/message"
)

type Facade struct {
	decodeMorse      *decoders.Morse
	decodeBinario    *decoders.Binario
	decodeMurcielago *decoders.Murcielago
}

func New() Facade {
	return Facade{
		decodeMorse:      &decoders.Morse{},
		decodeBinario:    &decoders.Binario{},
		decodeMurcielago: &decoders.Murcielago{},
	}
}
func (f *Facade) Encoders(name string) {
	m := message.Message{}
	m.AddObserver("morse", f.decodeMorse)
	m.AddObserver("binario", f.decodeBinario)
	m.AddObserver("murcielago", f.decodeMurcielago)
	m.Msg = name
	m.NotifyObservers()

}
