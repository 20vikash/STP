package sw

import (
	"STP/STP/consts"
	"STP/STP/models"
)

// Create and Add switches and interfaces

func CreateInterface(sw *models.Switch, type_ string) {
	interface_ := new(models.Interface)
	for i := range 3 {
		interface_.Type = type_
		interface_.Pair = nil
		interface_.Number = i

		sw.Interface = append(sw.Interface, interface_)
	}
}

func CreateSwitch() *models.Switch {
	// Creates a switch with fixed interfaces (3 Fast ethernets and 3 Gig Ethernets)
	// Create the interfaces.

	// Switch
	sw := new(models.Switch)

	// Fast ethernet
	CreateInterface(sw, consts.FASTETHERNET)

	// Gig Ethernet
	CreateInterface(sw, consts.GIGETHERNET)

	return sw
}
