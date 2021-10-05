package factory

import (
	".main.go/assemblyspot"
	veh ".main.go/vehicle"
	"fmt"
)

const assemblySpots int = 5

type Factory struct {
	AssemblingSpots chan *assemblyspot.AssemblySpot
}

func New() *Factory {
	factory := &Factory{
		AssemblingSpots: make(chan *assemblyspot.AssemblySpot, assemblySpots),
	}

	totalAssemblySpots := 0

	for {
		factory.AssemblingSpots <- &assemblyspot.AssemblySpot{}

		totalAssemblySpots++

		if totalAssemblySpots >= assemblySpots {
			break
		}
	}

	return factory
}

//HINT: this function is currently not returning anything, make it return right away every single vehicle once assembled,
//(Do not wait for all of them to be assembled to return them all, send each one ready over to main)
func (f *Factory) StartAssemblingProcess(amountOfVehicles int) chan *veh.Car {
	vehicleList := f.generateVehicleLots(amountOfVehicles)
	carChannel := make(chan *veh.Car, assemblySpots)

	for i, vehicle := range vehicleList {
		go func(v veh.Car) {
			fmt.Println("Assembling vehicle...", i)

			idleSpot := <-f.AssemblingSpots
			idleSpot.SetVehicle(&v)
			vehicle, err := idleSpot.AssembleVehicle()

			if err != nil {
				fmt.Println("there was an error while creating one vehicle")
				return
			}
			vehicle.TestingLog = f.testCar(vehicle)
			vehicle.AssembleLog = idleSpot.GetAssembledLogs()

			idleSpot.SetVehicle(nil)
			f.AssemblingSpots <- idleSpot
			carChannel <- vehicle
		}(vehicle)

	}
	return carChannel
}

func (Factory) generateVehicleLots(amountOfVehicles int) []veh.Car {
	var vehicles = []veh.Car{}
	var index = 0

	for {
		vehicles = append(vehicles, veh.Car{
			Id:            index,
			Chassis:       "NotSet",
			Tires:         "NotSet",
			Engine:        "NotSet",
			Electronics:   "NotSet",
			Dash:          "NotSet",
			Sits:          "NotSet",
			Windows:       "NotSet",
			EngineStarted: false,
		})

		index++

		if index >= amountOfVehicles {
			break
		}
	}

	return vehicles
}

func (f *Factory) testCar(car *veh.Car) string {
	logs := ""

	log, err := car.StartEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnLeft()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnRight()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.StopEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	return logs
}
