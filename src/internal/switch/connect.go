package sw

import "STP/STP/models"

func ConnectInterface(sw1 *models.Switch, sw2 *models.Switch, type_ string, sPort int, dPort int) {
	sw1.Interfaces[sPort].Pair = sw2.Interfaces[dPort]
	sw2.Interfaces[dPort].Pair = sw1.Interfaces[sPort]

	// Trigger STP operation once it is done
}
