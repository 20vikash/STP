package sw

import "STP/STP/models"

func ConnectInterface(sw1 *models.Switch, sw2 *models.Switch, type_ string, sPort int, dPort int) {
	// Switch interfaces are connected to neighbour switches
	sw1.Interfaces[sPort].Sw = sw2
	sw2.Interfaces[dPort].Sw = sw1

	sw1.Interfaces[sPort].Pair = sw2.Interfaces[dPort]
	sw2.Interfaces[dPort].Pair = sw1.Interfaces[sPort]

	// Trigger STP operation once it is done
}
