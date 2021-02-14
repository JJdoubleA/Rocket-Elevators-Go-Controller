package main

import (
	"fmt"
	"math"
	"sort"
)

var elevatorID = 1
var floorRequestButtonID = 1
var columnID = 1
var callButtonID = 1

type Battery struct {
	ID                        int
	status                    string
	amountOfFloors            int
	amountOfColumns           int
	amountOfBasements         int
	amountOfElevatorPerColumn int
	columnsList               []Column
	floorRequestButtonsList   []FloorRequestButton
}

func newBattery(id int, status string, amountOfFloors int, amountOfColumns int, amountOfBasements int, amountOfElevatorPerColumn int) Battery {
	b := Battery{}
	b.ID = id
	b.status = status
	b.amountOfFloors = amountOfFloors
	b.amountOfColumns = amountOfColumns
	b.amountOfBasements = amountOfBasements
	b.columnsList = []Column{}
	b.floorRequestButtonsList = []FloorRequestButton{}

	if amountOfBasements > 0 {
		b.createFloorRequestButtons(amountOfBasements)
		b.createBasementColumn(b.amountOfBasements, amountOfElevatorPerColumn)
		amountOfColumns--
	}

	b.createFloorRequestButtons(amountOfFloors)
	b.createColumns(amountOfColumns, amountOfFloors, amountOfElevatorPerColumn)

	return b
}

func (b *Battery) createBasementColumn(amountOfBasements int, amountOfElevatorPerColumn int) {
	servedFloors := []int{}
	floor := -1

	for i := 0; i < amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor--
	}

	col := Column{columnID, "online", servedFloors, b.amountOfFloors, amountOfBasements, true, []Elevator{}, []CallButton{}}
	b.columnsList = append(b.columnsList, col)
	columnID++

}

func (b *Battery) createColumns(amountOfColumns int, amountOfFloors int, amountOfElevatorPerColumn int) {
	amountOfFloorsPerColumn := int(math.Ceil(float64(b.amountOfFloors / amountOfElevatorPerColumn)))
	floor := 1

	for i := 0; i < amountOfColumns; i++ {
		servedFloors := []int{}
		for v := 0; v < amountOfFloorsPerColumn; v++ {

			if floor <= b.amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}

		col := Column{columnID, "online", servedFloors, b.amountOfFloors, amountOfElevatorPerColumn, false, []Elevator{}, []CallButton{}} //Missing some values in here
		b.columnsList = append(b.columnsList, col)
		columnID++
	}
}

func (b *Battery) createFloorRequestButtons(amountOfFloors int) {
	buttonFloor := 1

	for i := 0; i < amountOfFloors; i++ {
		floorRequestButtons := FloorRequestButton{floorRequestButtonID, "off", buttonFloor, "up"}
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, floorRequestButtons)
		buttonFloor++
		floorRequestButtonID++
	}
}

func (b *Battery) createBasementFloorRequestButtons(amountOfBasements int) {
	buttonFloor := -1

	for i := 0; i < amountOfBasements; i++ {
		floorRequestButtons := FloorRequestButton{floorRequestButtonID, "off", buttonFloor, "down"}
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, floorRequestButtons)
		buttonFloor--
		floorRequestButtonID++
	}
}

type Column struct {
	ID                int
	status            string
	servedFloors      []int
	amountOfFloors    int
	amountOfElevators int
	isBasement        bool
	elevatorsList     []Elevator
	callButtonsList   []CallButton
}

func (c *Column) createElevators(amountOfFloors int, amountOfElevators int) {

	for i := 0; i < amountOfElevators; i++ {
		elevator := Elevator{ID: elevatorID, status: "idle", amountOfFloors: amountOfFloors, direction: "null", currentFloor: 0, floorRequestList: []int{}, door: []Door{}}
		c.elevatorsList = append(c.elevatorsList, elevator)
		elevatorID++
	}
}

func (c *Column) createCallButtons(amountOfFloors int, isBasement bool) {
	if isBasement {
		buttonFloor := -1

		for i := 0; i < amountOfFloors; i++ {
			callButton := CallButton{callButtonID, "off", buttonFloor, "up"}
			c.callButtonsList = append(c.callButtonsList, callButton)
			buttonFloor--
			callButtonID++
		}
	} else {

		buttonFloor := 1

		for i := 0; i < amountOfFloors; i++ {
			callButton := CallButton{callButtonID, "off", buttonFloor, "down"}
			c.callButtonsList = append(c.callButtonsList, callButton)
			buttonFloor++
			callButtonID++
		}
	}
}

func (c Column) requestElevators(requestedFloor int, direction string) {
	elevator := c.findElevator(requestedFloor, direction)
	elevator.floorRequestList = append(elevator.floorRequestList, requestedFloor)

	elevator.sortFloorList()
	elevator.move()
	elevator.operateDoors()
}

func (c Column) findElevator(requestedFloor int, requestedDirection string) Elevator {

	bestElevator := Elevator{}
	bestScore := 5
	referenceGap := 1000000

	for _, elevator := range c.elevatorsList {
		if requestedFloor == elevator.currentFloor && elevator.status == "stopped" && requestedDirection == elevator.direction {
			bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, requestedFloor, bestElevator, bestScore, referenceGap)

		} else if requestedFloor == elevator.currentFloor && elevator.status == "up" && requestedDirection == elevator.direction {
			bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, requestedFloor, bestElevator, bestScore, referenceGap)

		} else if requestedFloor == elevator.currentFloor && elevator.status == "down" && requestedDirection == elevator.direction {
			bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, requestedFloor, bestElevator, bestScore, referenceGap)

		} else if elevator.status == "idle" {
			bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, requestedFloor, bestElevator, bestScore, referenceGap)

		} else {
			bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, requestedFloor, bestElevator, bestScore, referenceGap)
		}
	}
	return bestElevator
}

func (c Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator Elevator, requestedFloor int, bestElevator Elevator, bestScore int, referenceGap int) (Elevator, int, int) {

	if scoreToCheck < bestScore {
		bestScore = scoreToCheck
		bestElevator = newElevator
		referenceGap = int(math.Abs(float64(newElevator.currentFloor - requestedFloor)))

	} else if scoreToCheck == bestScore {

		gap := 0
		if newElevator.status == "idle" || newElevator.status == "stopped" {
			gap = int(math.Abs(float64(newElevator.currentFloor - requestedFloor)))
		} else {
			gap = int(math.Abs(float64(newElevator.currentFloor - newElevator.floorRequestList[0])))
		}

		if referenceGap > gap {
			bestElevator = newElevator
			referenceGap = gap
		}
	}
	return bestElevator, bestScore, referenceGap
}

type Elevator struct {
	ID               int
	status           string
	amountOfFloors   int
	direction        string
	currentFloor     int
	floorRequestList []int
	door             []Door
}

func (e *Elevator) requestFloor(requestedFloor int) {
	e.floorRequestList = append(e.floorRequestList, requestedFloor)
	e.sortFloorList()
	e.move()
	e.operateDoors()
}

func (e *Elevator) move() {
	for len(e.floorRequestList) != 0 {
		destination := e.floorRequestList[0]
		e.status = "moving"
		if e.currentFloor < destination {
			e.direction = "up"
			for e.currentFloor < destination {
				screenDisplay := e.currentFloor
				fmt.Println("CurrentFloor is", +screenDisplay)
				e.currentFloor++
			}
		} else if e.currentFloor > destination {
			e.direction = "down"
			for e.currentFloor > destination {
				screenDisplay := e.currentFloor
				fmt.Println("CurrentFloor is", +screenDisplay)
				e.currentFloor--
			}
		}
		e.status = "stopped"
		e.floorRequestList = e.floorRequestList[1:]
		fmt.Println("Elevator has stopped on floor: ", +e.currentFloor)
	}
}

func (e Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Slice(e.floorRequestList, func(i, j int) bool { return e.floorRequestList[i] < e.floorRequestList[j] })
	} else {
		sort.Slice(e.floorRequestList, func(i, j int) bool { return e.floorRequestList[i] > e.floorRequestList[j] })
	}
}

func (e *Elevator) operateDoors() {
	d := Door{}
	d.status = "opened"
	fmt.Println("Door is opening")
	d.status = "closed"
	fmt.Println("Door is closing")

}

type CallButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

type FloorRequestButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

type Door struct {
	ID     int
	status string
}

func main() {

	var servedFloors = []int{}

	battery := Battery{
		ID:                        1,
		amountOfColumns:           4,
		status:                    "online",
		amountOfFloors:            60,
		amountOfBasements:         6,
		amountOfElevatorPerColumn: 5,
		columnsList:               []Column{},
		floorRequestButtonsList:   []FloorRequestButton{},
	}

	battery.createBasementColumn(battery.amountOfBasements, battery.amountOfElevatorPerColumn)
	battery.createColumns(battery.amountOfColumns, battery.amountOfFloors, battery.amountOfElevatorPerColumn)
	battery.createFloorRequestButtons(battery.amountOfFloors)
	battery.createBasementFloorRequestButtons(battery.amountOfBasements)

	column := Column{
		ID:                columnID,
		status:            "online",
		servedFloors:      servedFloors,
		amountOfFloors:    60,
		amountOfElevators: 5,
		elevatorsList:     []Elevator{},
		callButtonsList:   []CallButton{},
	}

	column.createElevators(column.amountOfFloors, column.amountOfElevators)
	column.createCallButtons(60, false)

	//Sceanario 1

	elevator := Elevator{
		ID:               elevatorID,
		status:           "idle",
		amountOfFloors:   60,
		direction:        "nill",
		currentFloor:     1,
		floorRequestList: []int{},
		door:             []Door{},
	}

	//fmt.Println("Elevator B1")
	column.elevatorsList[0].currentFloor = 20
	column.elevatorsList[0].status = "down"
	column.elevatorsList[0].direction = "down"
	column.elevatorsList[0].floorRequestList = append(column.elevatorsList[0].floorRequestList, 5)

	//fmt.Println("ElevatorB2")
	column.elevatorsList[1].currentFloor = 3
	column.elevatorsList[1].status = "up"
	column.elevatorsList[1].direction = "up"
	column.elevatorsList[1].floorRequestList = append(column.elevatorsList[1].floorRequestList, 15)

	//fmt.Println("ElevatorB3")
	column.elevatorsList[2].currentFloor = 13
	column.elevatorsList[2].status = "down"
	column.elevatorsList[2].direction = "down"
	column.elevatorsList[2].floorRequestList = append(column.elevatorsList[2].floorRequestList, 1)

	//fmt.Println("ElevatorB4")
	column.elevatorsList[3].currentFloor = 15
	column.elevatorsList[3].status = "down"
	column.elevatorsList[3].direction = "down"
	column.elevatorsList[3].floorRequestList = append(column.elevatorsList[3].floorRequestList, 2)

	//fmt.Println("ElevatorB5")
	column.elevatorsList[4].currentFloor = 6
	column.elevatorsList[4].status = "down"
	column.elevatorsList[4].direction = "down"
	column.elevatorsList[4].floorRequestList = append(column.elevatorsList[4].floorRequestList, 1)

	column.findElevator(1, "up")
	column.requestElevators(1, "up")
	elevator.requestFloor(20)
	//fmt.Printf("%v", column.elevatorsList)

	//Scenario 2
	/*
		elevator := Elevator{
			ID:               elevatorID,
			status:           "idle",
			amountOfFloors:   60,
			direction:        "nill",
			currentFloor:     20,
			floorRequestList: []int{},
			door:             []Door{},
		}

		//fmt.Println("Elevator C1")
		column.elevatorsList[0].currentFloor = 20
		column.elevatorsList[0].status = "idle"
		column.elevatorsList[0].direction = "up"
		column.elevatorsList[0].floorRequestList = append(column.elevatorsList[0].floorRequestList, 21)

		//fmt.Println("Elevator C2")
		column.elevatorsList[1].currentFloor = 23
		column.elevatorsList[1].status = "up"
		column.elevatorsList[1].direction = "up"
		column.elevatorsList[1].floorRequestList = append(column.elevatorsList[1].floorRequestList, 28)

		//fmt.Println("Elevator C3")
		column.elevatorsList[2].currentFloor = 33
		column.elevatorsList[2].status = "down"
		column.elevatorsList[2].direction = "down"
		column.elevatorsList[2].floorRequestList = append(column.elevatorsList[2].floorRequestList, 20)

		//fmt.Println("Elevator C4")
		column.elevatorsList[3].currentFloor = 40
		column.elevatorsList[3].status = "down"
		column.elevatorsList[3].direction = "down"
		column.elevatorsList[3].floorRequestList = append(column.elevatorsList[3].floorRequestList, 24)

		//fmt.Println("Elevator C5")
		column.elevatorsList[4].currentFloor = 39
		column.elevatorsList[4].status = "down"
		column.elevatorsList[4].direction = "down"
		column.elevatorsList[4].floorRequestList = append(column.elevatorsList[4].floorRequestList, 20)

		column.findElevator(20, "up")
		column.requestElevators(20, "up")
		elevator.requestFloor(36)
		//fmt.Printf("%v", column.elevatorsList)
	*/
	//Scenario 3
	/*
		//fmt.Println("Elevator D1")
		column.elevatorsList[0].currentFloor = 58
		column.elevatorsList[0].status = "moving"
		column.elevatorsList[0].direction = "down"
		column.elevatorsList[0].floorRequestList = append(column.elevatorsList[0].floorRequestList, 40)

		//fmt.Println("Elevator D2")
		column.elevatorsList[1].currentFloor = 50
		column.elevatorsList[1].status = "up"
		column.elevatorsList[1].direction = "up"
		column.elevatorsList[1].floorRequestList = append(column.elevatorsList[1].floorRequestList, 60)

		//fmt.Println("Elevator D3")
		column.elevatorsList[2].currentFloor = 46
		column.elevatorsList[2].status = "up"
		column.elevatorsList[2].direction = "up"
		column.elevatorsList[2].floorRequestList = append(column.elevatorsList[2].floorRequestList, 58)

		//fmt.Println("Elevator D4")
		column.elevatorsList[3].currentFloor = 40
		column.elevatorsList[3].status = "up"
		column.elevatorsList[3].direction = "up"
		column.elevatorsList[3].floorRequestList = append(column.elevatorsList[3].floorRequestList, 54)

		//fmt.Println("Elevator D5")
		column.elevatorsList[4].currentFloor = 60
		column.elevatorsList[4].status = "down"
		column.elevatorsList[4].direction = "down"
		column.elevatorsList[4].floorRequestList = append(column.elevatorsList[4].floorRequestList, 40)

		column.findElevator(54, "down")
		column.requestElevators(54, "down")
		//fmt.Printf("%v", column.elevatorsList)
	*/
}
