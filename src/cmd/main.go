package main

import (
	"STP/STP/models"
	"fmt"
)

var sws = make([]*models.Switch, 0)
var connections = make(map[*models.Interface]*models.Interface)

// var rootBridge = new(models.Switch)
var bpduChan chan *models.BPDU = make(chan *models.BPDU, 10)

func main() {
	fmt.Println("STP simulator")
	fmt.Println("1. Create switch\n 2. Connect interface\n -1. Quit")

	for {
		fmt.Println("Enter your choice: ")

		var choice int

		fmt.Scan(&choice)

		if choice == -1 {
			break
		}

		switch choice {
		case 1:
			{
				sw := CreateSwitch()
				sws = append(sws, sw)
			}
		case 2:
			{
				ConnectInterface()
			}
		}
	}
}
