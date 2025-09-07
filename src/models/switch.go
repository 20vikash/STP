package models

type Switch struct {
	Interface Interface
}

type Interface struct {
	Fe1 FastEthernet
	Fe2 FastEthernet
	Fe3 FastEthernet
	Ge1 GigEthernet
	Ge2 GigEthernet
	Ge3 GigEthernet
}
