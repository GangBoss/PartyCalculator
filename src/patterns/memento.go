package patterns

type Memento struct {
	memento interface{}
}

func (m *Memento) SetMemento(i interface{}) {
	m.memento = i
}

func (m *Memento) GetMemento() interface{} {
	return m.memento
}
