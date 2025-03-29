package main

import (
	"fmt"
)

func main() {
	// Input required values
	var age, intensity, rhr float64

	fmt.Print("How old are you?: ")
	fmt.Scan(&age)

	fmt.Print("What intensity are you training(%)?: ")
	fmt.Scan(&intensity)

	fmt.Print("What is your resting heart rate(bpm)?: ")
	fmt.Scan(&rhr)

	// Max heart rate
	mhr := maxHeartRate(age)

	// Basic target heart rate
	intensityFloat := intensityConvert(intensity)
	thrBasic := basicTHR(mhr, intensityFloat)

	// Target heart rate using the Karvonen formula
	thrKarvonen := karvonenTHR(mhr, intensityFloat, rhr)

	// Outputs
	fmt.Printf("Your target heart rate (basic): %.0fbpm\n", thrBasic)
	fmt.Printf("Your target heart rate (Karvonen): %.0fbpm\n", thrKarvonen)

}

func maxHeartRate(age float64) (mhr float64) {
	mhr = 220 - age
	return mhr
}

func basicTHR(mhr, intensity float64) (thr float64) {
	thr = mhr * intensity
	return thr
}

func karvonenTHR(mhr, intensity, rhr float64) (thr float64) {
	thr = ((mhr - rhr) * intensity) + rhr
	return thr
}

func intensityConvert(intensity float64) (intensityFloat float64) {
	intensityFloat = intensity / 100
	return intensityFloat
}
