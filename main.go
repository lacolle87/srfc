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

func help() {
	helpMessage := fmt.Sprintf(`Usage: racer-calculator [OPTIONS]

Description:
  Calculates total fuel, total laps, and stints required based on race settings.
  Ideal for endurance racing, focusing on fuel consumption and tank capacity.

Options:
  -lt    Lap time in seconds
  -rl    Race length in minutes
  -fc    Fuel consumption per lap
  -tc    Fuel tank capacity
  -ef    Extra fuel needed in laps

Examples:
  racer-calculator -lt 90 -rl 120 -fc 3.5 -tc 110

Output:
  The output displays the total fuel, total laps, and stints required based on the provided race settings.`,
	)
	fmt.Println(helpMessage)
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
	flag.Float64Var(&raceData.ExtraFuel, "ef", 2, "Extra fuel needed in laps")

	showHelp := flag.Bool("h", false, "Show help message")
	flag.Parse()

	if *showHelp {
		help()
		return
	}

	if raceData.LapSec == 0 || raceData.RaceMin == 0 || raceData.FuelPerLap == 0 {
		fmt.Println("Some required flags are missing. Entering interactive mode...")
		raceData.promptUserInput()
	}

	raceData.calculateLaps()
	raceData.calculateFuel()
	raceData.calculateStints()
	if raceData.ExtraFuel != 0 {
		raceData.addExtraFuel()
	}
	fmt.Printf("Total fuel: %.2f\nTotal laps: %d\nStints: %d\nAdded extra fuel for %.1f laps\n", raceData.TotalFuel, raceData.TotalLaps, raceData.Stints, raceData.ExtraFuel)
}
