package main

import (
	"flag"
	"fmt"
	"math"
)

type RaceData struct {
	LapTimeInSec          uint
	RaceLengthInMin       uint
	Stints                uint
	TotalLaps             uint
	FuelConsumptionPerLap float64
	FuelTankCapacity      float64
	TotalFuel             float64
	ExtraFuel             float64
}

func (r *RaceData) calcLaps() {
	r.TotalLaps = r.RaceLengthInMin * 60 / r.LapTimeInSec
}

func (r *RaceData) addExtraFuel() {
	r.TotalFuel += r.FuelConsumptionPerLap * r.ExtraFuel
}

func (r *RaceData) calcFuel() {
	r.TotalFuel = float64(r.TotalLaps) * r.FuelConsumptionPerLap
}

func (r *RaceData) calcStints() {
	r.Stints = uint(math.Ceil(r.TotalFuel / r.FuelTankCapacity))
}

func main() {
	raceData := &RaceData{}

	flag.UintVar(&raceData.LapTimeInSec, "lap-time", 0, "Time for one lap in seconds")
	flag.UintVar(&raceData.RaceLengthInMin, "race-length", 0, "Race length in minutes")
	flag.Float64Var(&raceData.FuelConsumptionPerLap, "fuel-consumption", 0, "Fuel consumption per lap")
	flag.Float64Var(&raceData.FuelTankCapacity, "fuel-tank", 0, "Fuel tank capacity")
	flag.Float64Var(&raceData.ExtraFuel, "extra-fuel", 0, "Extra fuel allowance")

	flag.Parse()

	if raceData.LapTimeInSec == 0 || raceData.RaceLengthInMin == 0 || raceData.FuelConsumptionPerLap == 0 || raceData.FuelTankCapacity == 0 {
		fmt.Println("Some required flags are missing. Entering interactive mode...")
		fmt.Print("Enter lap time (in seconds): ")
		fmt.Scan(&raceData.LapTimeInSec)
		fmt.Print("Enter race length (in minutes): ")
		fmt.Scan(&raceData.RaceLengthInMin)
		fmt.Print("Enter fuel consumption per lap: ")
		fmt.Scan(&raceData.FuelConsumptionPerLap)
		fmt.Print("Enter fuel tank capacity: ")
		fmt.Scan(&raceData.FuelTankCapacity)
		fmt.Print("Enter extra fuel allowance: ")
		fmt.Scan(&raceData.ExtraFuel)
	}

	raceData.calcLaps()
	raceData.calcFuel()
	raceData.calcStints()
	raceData.addExtraFuel()

	fmt.Printf("Total fuel: %.2f\nTotal laps: %d\nStints: %d\n", raceData.TotalFuel, raceData.TotalLaps, raceData.Stints)
}
