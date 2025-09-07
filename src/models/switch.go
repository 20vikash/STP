package models

type Switch struct {
	Id        int
	Interface []*Interface
}

type Interface struct {
	Type   string // FastEthernet and GigEthernet consts
	Number int
	Pair   *Switch
}
