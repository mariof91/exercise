package main

import (
	".main.go/factory"
	"fmt"
)

const carsAmount = 100

func main() {
	factory := factory.New()

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id
	cars := factory.StartAssemblingProcess(carsAmount)

	for i := 0; i < carsAmount; i++ {
		car := <-cars
		fmt.Println("Inside main, car ID: ", car.Id)
		fmt.Println("Inside main, assemble log:", car.AssembleLog)
		fmt.Println("Inside main, testLog: ", car.TestingLog)

	}
}