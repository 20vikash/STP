package main

import (
	"STP/STP/consts"
	"STP/STP/models"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func CreateSwitch() *models.Switch {
	sw := models.CreateSwitch()

	fmt.Println("Switch successfully created with 3 fastEthernet and 3 GigEthernet")
	fmt.Printf("Mac address is: %s", sw.MacAddr)

	go InitiateHelloBPDU(sw) // Each goRoutine for every switches

	return sw
}

func findInterfaces(Switch int, Port string) *models.Interface {
	var type_ string

	for _, sw := range sws {
		if sw.Id == Switch {
			for _, p := range sw.Interfaces {
				if string(Port[0]) == "f" {
					type_ = consts.FASTETHERNET
				} else if string(Port[0]) == "g" {
					type_ = consts.GIGETHERNET
				}
				num, _ := strconv.Atoi(string(Port[1]))
				if p.Type == type_ && p.Number == num {
					return p
				}
			}
		}
	}

	return nil
}

func ConnectInterface() error {
	var sSwitch int
	var dSwitch int
	var sPort string
	var dPort string

	var count int = 0

	fmt.Println("The available switches: ")

	for i, sw := range sws {
		fmt.Printf("%d: SW %d\n", i+1, sw.Id)
	}

	// Eg: 1, 2
	fmt.Print("Enter the source switch: ")
	fmt.Scan(&sSwitch)

	fmt.Print("Enter the destination switch: ")
	fmt.Scan(&dSwitch)

	// Eg: f0, g1, g2
	fmt.Println("Enter the source port: ")
	fmt.Scan(&sPort)

	fmt.Println("Enter the destination port: ")
	fmt.Scan(&dPort)

	for _, sw := range sws {
		if sw.Id == sSwitch {
			count += 1
		}
		if sw.Id == dSwitch {
			count += 1
		}
	}

	if count != 2 {
		fmt.Println("Source or destination switch dosent exist")
		return errors.New("not found")
	}

	if sSwitch == dSwitch {
		fmt.Println("Source and destinaion switch cannot be the same")
		return errors.New("conflict")
	}

	// Source interface
	sInterface := findInterfaces(sSwitch, sPort)

	// Destination interface
	dInterface := findInterfaces(dSwitch, dPort)

	if sInterface == nil || dInterface == nil {
		return errors.New("no interface")
	}

	// Check if the pair already exists
	if _, ok := connections[sInterface]; ok {
		fmt.Println("Port not free")
		return errors.New("conflict")
	}
	if _, ok := connections[dInterface]; ok {
		fmt.Println("Port not free")
		return errors.New("conflict")
	}

	connections[sInterface] = dInterface
	connections[dInterface] = sInterface

	fmt.Println("Connection successfully made")

	// Reset all timers
	for _, sw := range sws {
		sw.TimerResetChan <- true
	}

	return nil
}

func InitiateHelloBPDU(sw *models.Switch) {
	isRoot := true
	bpduTimer := time.NewTicker(2 * time.Second)
	timeOut := time.NewTimer(6 * time.Second)

	for {
		select {
		case <-sw.TimerResetChan:
			timeOut.Reset(6 * time.Second)
		case <-bpduTimer.C:
			if isRoot { // Send hello BPDU frames every 2 seconds if its the root bridge
				for key, value := range connections {
					if key.Sw == sw {
						bpdu := models.CreateBPDU(key, value, key.Sw.MacAddr)
						value.Sw.BpduChan <- bpdu
					}
				}
			}
		case bpdu := <-sw.BpduChan: // Recieve BPDU frames from the neighbour bridge
			timeOut.Reset(6 * time.Second)
			for _, inter := range sw.Interfaces {
				if bpdu.DInterface == inter {
					if inter.Priority == bpdu.BridgeId {
						if sw.MacAddr > bpdu.MacAddr {
							isRoot = false
							ForwardBPDU(inter, bpdu, bpdu.MacAddr)
						}
						break
					} else if inter.Priority > bpdu.BridgeId {
						isRoot = false
						ForwardBPDU(inter, bpdu, bpdu.MacAddr)
						break
					}
				}
			}
		case <-timeOut.C:
			isRoot = true
			fmt.Printf("The new root bridge is %d\n", sw.Id)
		}
	}
}

func ForwardBPDU(inter *models.Interface, bpdu *models.BPDU, mac string) {
	for key, value := range connections {
		if value != bpdu.SInterface && key.Sw == inter.Sw {
			bpdu := models.CreateBPDU(inter, value, mac)
			value.Sw.BpduChan <- bpdu
		}
	}
}
