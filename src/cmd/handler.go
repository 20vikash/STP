package main

import (
	"STP/STP/consts"
	"STP/STP/models"
	"errors"
	"fmt"
	"strconv"
)

func CreateSwitch() *models.Switch {
	sw := models.CreateSwitch()

	fmt.Println("Switch successfully created with 3 fastEthernet and 3 GigEthernet")

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

	connections[sInterface] = dInterface
	connections[dInterface] = sInterface

	fmt.Println("Connection successfully made")

	return nil
}
