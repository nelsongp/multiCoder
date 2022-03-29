package observer

type Observer interface {
	Notify(name string)
}
