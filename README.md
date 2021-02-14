Comment and Uncomment Each scenario for the Csharp program

//Scenario 1
             
               column.elevatorList[0].currentFloor = 20;
               column.elevatorList[0].status = "moving";
               column.elevatorList[0].direction = "down";
               column.elevatorList[0].floorRequestList.Add(5);
               
               column.elevatorList[1].currentFloor = 3;
               column.elevatorList[1].status = "moving";
               column.elevatorList[1].direction = "up";
               column.elevatorList[0].floorRequestList.Add(15);

               column.elevatorList[2].currentFloor = 13;
               column.elevatorList[2].status = "moving";
               column.elevatorList[2].direction = "down";
               column.elevatorList[0].floorRequestList.Add(1);

               column.elevatorList[3].currentFloor = 15;
               column.elevatorList[3].status = "moving";
               column.elevatorList[3].direction = "down";
               column.elevatorList[0].floorRequestList.Add(2);

               column.elevatorList[4].currentFloor = 6;
               column.elevatorList[4].status = "moving";
               column.elevatorList[4].direction = "down";
               column.elevatorList[0].floorRequestList.Add(1);
 
               elevator = column.requestElevator(1,"up");
               elevator.requestFloor(20);

Comment and Uncomment each scenario for the GO main

//Scenario 2

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

