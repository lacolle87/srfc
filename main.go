package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"strconv"
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

func (r *RaceData) validateDigits(input string) error {
	if val, err := strconv.Atoi(input); err != nil || val <= 0 {
		return errors.New("input must contain only digits")
	}
	return nil
}

func (r *RaceData) validateFloat(input string) error {
	if val, err := strconv.ParseFloat(input, 64); err != nil || val <= 0 {
		return errors.New("input must be a positive floating-point number")
	}
	return nil
}

func (r *RaceData) validateExtraFuelFloat(input string) error {
	if val, err := strconv.ParseFloat(input, 64); err != nil || val < 0 {
		return errors.New("input must be a positive floating-point number")
	}
	return nil
}

func (r *RaceData) promptUserInput() {
	r.promptInput("Enter lap time (in seconds): ", &r.LapSec, r.validateDigits)
	r.promptInput("Enter race length (in minutes): ", &r.RaceMin, r.validateDigits)
	r.promptInput("Enter fuel consumption per lap: ", &r.FuelPerLap, r.validateFloat)
	r.promptInput("Enter fuel tank capacity: ", &r.FuelCap, r.validateFloat)
	r.promptInput("Enter extra fuel needed in laps: ", &r.ExtraFuel, r.validateExtraFuelFloat)
}

func (r *RaceData) promptInput(prompt string, target interface{}, validator func(string) error) {
	for {
		var input string
		fmt.Print(prompt)
		fmt.Scan(&input)
		if err := validator(input); err == nil {
			switch v := target.(type) {
			case *uint:
				val, _ := strconv.Atoi(input)
				*v = uint(val)
			case *float64:
				*v, _ = strconv.ParseFloat(input, 64)
			}
			break
		}
		fmt.Println("Invalid input, please try again.")
	}
}

func help() {
	fmt.Println(`Usage: racer-calculator [OPTIONS]

Description:
  Calculates total fuel, total laps, and stints required based on race settings.
  Ideal for endurance racing, focusing on fuel consumption and tank capacity.

Options:
  -lt    Lap time in seconds (digits only)
  -rl    Race length in minutes (digits only)
  -fc    Fuel consumption per lap
  -tc    Fuel tank capacity
  -ef    Extra fuel needed in laps

Examples:
  racer-calculator -lt 90 -rl 120 -fc 3.5 -tc 110`)
}

func main() {
	raceData := &RaceData{}

	flag.UintVar(&raceData.LapSec, "lt", 0, "Lap time in seconds (digits only)")
	flag.UintVar(&raceData.RaceMin, "rl", 0, "Race length in minutes (digits only)")
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
		fmt.Println("Some required flags are missing or invalid. Entering interactive mode...")
		raceData.promptUserInput()
	}

	raceData.calculateLaps()
	raceData.calculateFuel()
	raceData.calculateStints()
	if raceData.ExtraFuel != 0 {
		raceData.addExtraFuel()
	}

	fmt.Printf("Total fuel: %.2f\nTotal laps: %d\nStints: %d\nAdded extra fuel for %.1f laps\n",
		raceData.TotalFuel, raceData.TotalLaps, raceData.Stints, raceData.ExtraFuel)
}
