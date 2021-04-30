package encoder
type Encoder interface{
	Morse(string) string
	Binario(string) string
	Murcielag(string) string
}
