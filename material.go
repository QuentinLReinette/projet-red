package main

type materialType int

const (
	wolfHide materialType = iota
	trollHide
	boarLeather
	crowFeather
)

var materialNames = map[materialType]string{
	wolfHide:    "Wolf Hide",
	trollHide:   "Troll Hide",
	boarLeather: "Boar Leather",
	crowFeather: "Crow Feather",
}

type material struct {
	materialType
	quantity int
}

func (m material) String() string {
	return materialNames[m.materialType]
}

func (m materialType) String() string {
	return materialNames[m]
}

func (m material) use() {
	println("This item does nothing on its own.")
}
