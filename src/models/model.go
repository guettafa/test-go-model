package model

type Model interface {
	GetName() string
	SetName(newName string)
}

type modelImpl struct {
	name string
}

func (m modelImpl) GetName() string {
	return m.name
}
	
func (m *modelImpl) SetName(newName string) {
	m.name = newName
}
