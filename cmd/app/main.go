package main

import (
	"fmt"
	"math"
)

type RaceData struct {
	LapTimeInSec          uint16
	RaceLengthInMin       uint16
	Stints                uint16
	TotalLaps             uint16
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
	r.Stints = uint16(math.Ceil(r.TotalFuel / r.FuelTankCapacity))
}

func main() {
	raceData := RaceData{
		LapTimeInSec:          60,
		RaceLengthInMin:       30,
		Stints:                0,
		TotalLaps:             0,
		FuelConsumptionPerLap: 2,
		FuelTankCapacity:      100,
		TotalFuel:             0,
		ExtraFuel:             2,
	}

	raceData.calcLaps()
	raceData.calcFuel()
	raceData.calcStints()
	raceData.addExtraFuel()
	fmt.Printf("Total fuel: %.2f\nTotal laps: %d\nStints: %d", raceData.TotalFuel, raceData.TotalLaps, raceData.Stints)
}
