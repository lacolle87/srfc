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

func (r *RaceData) calcLaps() {
	r.TotalLaps = r.RaceMin * 60 / r.LapSec
}

func (r *RaceData) addExtraFuel() {
	r.TotalFuel += r.FuelPerLap * r.ExtraFuel
}

func (r *RaceData) calcFuel() {
	r.TotalFuel = float64(r.TotalLaps) * r.FuelPerLap
}

func (r *RaceData) calcStints() {
	r.Stints = uint(math.Ceil(r.TotalFuel / r.FuelCap))
}

func main() {
	raceData := &RaceData{}

	flag.UintVar(&raceData.LapSec, "lt", 0, "Lap time in seconds")
	flag.UintVar(&raceData.RaceMin, "rl", 0, "Race length in minutes")
	flag.Float64Var(&raceData.FuelPerLap, "fc", 0, "Fuel consumption per lap")
	flag.Float64Var(&raceData.FuelCap, "tc", 0, "Fuel tank capacity")
	flag.Float64Var(&raceData.ExtraFuel, "ef", 0, "Extra fuel needed, default: 2")

	flag.Parse()

	if raceData.LapSec == 0 || raceData.RaceMin == 0 || raceData.FuelPerLap == 0 || raceData.FuelCap == 0 {
		fmt.Println("Some required flags are missing. Entering interactive mode...")
		fmt.Print("Enter lap time (in seconds): ")
		fmt.Scan(&raceData.LapSec)
		fmt.Print("Enter race length (in minutes): ")
		fmt.Scan(&raceData.RaceMin)
		fmt.Print("Enter fuel consumption per lap: ")
		fmt.Scan(&raceData.FuelPerLap)
		fmt.Print("Enter fuel tank capacity: ")
		fmt.Scan(&raceData.FuelCap)
		fmt.Print("Enter extra fuel needed in laps: ")
		fmt.Scan(&raceData.ExtraFuel)
	}

	raceData.calcLaps()
	raceData.calcFuel()
	raceData.calcStints()
	raceData.addExtraFuel()

	fmt.Printf("Total fuel: %.2f\nTotal laps: %d\nStints: %d\n", raceData.TotalFuel, raceData.TotalLaps, raceData.Stints)
}
