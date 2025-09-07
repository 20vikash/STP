package models

import "STP/STP/consts"

type Switch struct {
	Id         int
	Interfaces []*Interface
}

type Interface struct {
	Type   string // FastEthernet and GigEthernet consts
	Number int
	Pair   *Switch
}

var ID = 0

func CreateInterface(sw *Switch, type_ string) {
	for i := range 3 {
		interface_ := new(Interface)
		interface_.Type = type_
		interface_.Pair = nil
		interface_.Number = i

		sw.Interfaces = append(sw.Interfaces, interface_)
	}
}

func CreateSwitch() *Switch {
	// Creates a switch with fixed interfaces (3 Fast ethernets and 3 Gig Ethernets)
	// Create the interfaces.

	// Switch
	sw := new(Switch)

	// Fast ethernet
	CreateInterface(sw, consts.FASTETHERNET)

	// Gig Ethernet
	CreateInterface(sw, consts.GIGETHERNET)

	sw.Id = ID
	ID += 1 // For the next switch

	return sw
}
