package encoder
type Encoder interface{
	Morse(string) string
	Murcielago(string) string
	Binario(string) string
}
