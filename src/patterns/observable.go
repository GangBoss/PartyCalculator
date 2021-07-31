package patterns

// Observable not usable for golang use channels
type Observable interface {
	attach(Observer)
	detach(Observer)
	notify()
}

type Observer interface {
	update()
}
