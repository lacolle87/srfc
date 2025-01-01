package main

import (
	"flag"
	"fmt"
	"math"
)

type RaceData struct {
	LapSec     uint
	RaceMin    uint
	Stints     uint
	TotalLaps  uint
	FuelPerLap float64
	FuelCap    float64
	TotalFuel  float64
	ExtraFuel  float64
}

func (r *RaceData) calculateLaps() {
	r.TotalLaps = r.RaceMin * 60 / r.LapSec
}

func (r *RaceData) calculateFuel() {
	r.TotalFuel = float64(r.TotalLaps) * r.FuelPerLap
}

func (r *RaceData) calculateStints() {
	r.Stints = uint(math.Ceil(r.TotalFuel / r.FuelCap))
}

func (r *RaceData) addExtraFuel() {
	r.TotalFuel += r.FuelPerLap * r.ExtraFuel
}

func (r *RaceData) promptUserInput() {
	fmt.Print("Enter lap time (in seconds): ")
	fmt.Scan(&r.LapSec)
	fmt.Print("Enter race length (in minutes): ")
	fmt.Scan(&r.RaceMin)
	fmt.Print("Enter fuel consumption per lap: ")
	fmt.Scan(&r.FuelPerLap)
	fmt.Print("Enter fuel tank capacity: ")
	fmt.Scan(&r.FuelCap)
	fmt.Print("Enter extra fuel needed in laps: ")
	fmt.Scan(&r.ExtraFuel)
}

func main() {
	raceData := &RaceData{}

	flag.UintVar(&raceData.LapSec, "lt", 0, "Lap time in seconds")
	flag.UintVar(&raceData.RaceMin, "rl", 0, "Race length in minutes")
	flag.Float64Var(&raceData.FuelPerLap, "fc", 0, "Fuel consumption per lap")
	flag.Float64Var(&raceData.FuelCap, "tc", 110, "Fuel tank capacity")
	flag.Float64Var(&raceData.ExtraFuel, "ef", 2, "Extra fuel needed")

	flag.Parse()

	if raceData.LapSec == 0 || raceData.RaceMin == 0 || raceData.FuelPerLap == 0 {
		fmt.Println("Some required flags are missing. Entering interactive mode...")
		raceData.promptUserInput()
	}

	raceData.calculateLaps()
	raceData.calculateFuel()
	raceData.calculateStints()
	raceData.addExtraFuel()

	fmt.Printf("Total fuel: %.2f\nTotal laps: %d\nStints: %d\n", raceData.TotalFuel, raceData.TotalLaps, raceData.Stints)
}
