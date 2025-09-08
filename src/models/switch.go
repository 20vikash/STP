package models

import (
	"STP/STP/consts"
	"crypto/rand"
	"fmt"
)

type Switch struct {
	Id             int
	Interfaces     []*Interface
	MacAddr        string
	BpduChan       chan *BPDU
	TimerResetChan chan bool
}

type Interface struct {
	Type     string // FastEthernet and GigEthernet consts
	Number   int
	Priority int
	Sw       *Switch
	Pair     *Interface
}

var ID = 1

func CreateInterface(sw *Switch, type_ string) {
	for i := range 3 {
		interface_ := new(Interface)
		interface_.Type = type_
		interface_.Pair = nil
		interface_.Sw = sw
		interface_.Number = i

		sw.Interfaces = append(sw.Interfaces, interface_)
	}
}

func GenerateMac() string {
	mac := make([]byte, 6)

	_, err := rand.Read(mac)
	if err != nil {
		panic(err)
	}

	mac[0] = (mac[0] | 2) & 0xFE

	return fmt.Sprintf("%02X:%02X:%02X:%02X:%02X:%02X",
		mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])
}

func SetPriorityMac(sw *Switch) {
	for _, inter := range sw.Interfaces {
		inter.Priority = 327678
	}
}

func CreateSwitch() *Switch {
	// Creates a switch with fixed interfaces (3 Fast ethernets and 3 Gig Ethernets)
	// Create the interfaces.

	// Switch
	sw := new(Switch)
	sw.MacAddr = GenerateMac()
	sw.BpduChan = make(chan *BPDU)
	sw.TimerResetChan = make(chan bool)

	// Fast ethernet
	CreateInterface(sw, consts.FASTETHERNET)

	// Gig Ethernet
	CreateInterface(sw, consts.GIGETHERNET)

	SetPriorityMac(sw)

	sw.Id = ID
	ID += 1 // For the next switch

	return sw

	// Trigger STP operation here
}
