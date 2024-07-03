package main

import "math"

type Body struct {
	Mass float64
	PosX float64
	PosY float64
	VelX float64
	VelY float64
}

const GConst = 6.6743e-11

func getDistance(mainBody Body, otherBody Body) float64 {
	distX2 := math.Pow(2, mainBody.PosX-otherBody.PosX)
	distY2 := math.Pow(2, mainBody.PosY-otherBody.PosY)

	return math.Pow(0.5, distX2+distY2)
}

func calculateVdotX(mainBody Body, otherBody Body) float64 {
	xVdot := -(mainBody.PosX - otherBody.PosX) * (GConst * otherBody.Mass) / math.Pow(3, getDistance(mainBody, otherBody))
	return xVdot
}

func calculateVdotY(mainBody Body, otherBody Body) float64 {
	yVdot := -(mainBody.PosY - otherBody.PosY) * (GConst * otherBody.Mass) / math.Pow(3, getDistance(mainBody, otherBody))
	return yVdot
}

func updateVeloecity(mainBody Body, otherBody Body) {
	// Update the velocity of the body from the influence of another body.
	mainBody.VelX = mainBody.VelX + calculateVdotX(mainBody, otherBody)
	mainBody.VelY = mainBody.VelY + calculateVdotY(mainBody, otherBody)
	//Run this on with all Other Bodies to Get resultant force.
}

func updatePosition(mainBody Body, timeStep float64) {
	mainBody.PosX += mainBody.VelX * timeStep
	mainBody.PosY += mainBody.VelY * timeStep
}

func main() {
	timeStep := 1

	numberOfTimeSteps := 365 * 60 * 60 * 24
	checkInTime := 60 * 60 * 24

	Sun := Body{2e30, 0., 0., 0., 0.}
	Mercury := Body{3.3e23, 0, 57e9, 47.3e3, 0}
	Venus := Body{4.8e24, 0, 108e9, 35e3, 0}
	Earth := Body{6e24, 0, -150e9, -29782, 0}
	Mars := Body{6.4e23, 0, 227e9, 24e3, 0}

	var allBodies [5]Body
	allBodies[0] = Sun
	allBodies[1] = Mercury
	allBodies[2] = Venus
	allBodies[3] = Earth
	allBodies[4] = Mars

	for nT := 0; nT < numberOfTimeSteps; nT++ {
		for i := 0; i < len(allBodies); i++ {
			for j := 0; j < len(allBodies); j++ {
				if i == j {
					continue
				} else {
					updateVeloecity(allBodies[i], allBodies[j])
				}
			}
			updatePosition(allBodies[i], 1.0)
		}
		if nT%checkInTime == 0 {
			println("Time Passed : ", timeStep*nT/checkInTime, " Days")
		}

	}

}
