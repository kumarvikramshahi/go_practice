package main

import (
	"fmt"
)

type motorMilage interface {
	milage() float64
	totalTimeRun() float64
}

type BMW struct {
	fuel          float64
	avgSpeed      float64
	distTravelled float64
}

type Audi struct {
	fuel          float64
	distTravelled float64
}

func (c BMW) milage() float64 {
	fmt.Println(c.distTravelled, "/", c.fuel)
	return c.distTravelled / c.fuel
}

func (c Audi) milage() float64 {
	return c.distTravelled / c.fuel
}

func (c BMW) totalTimeRun() float64 {
	return c.distTravelled / c.avgSpeed
}

func Interface() {
	fmt.Println("hi")
	c1 := BMW{
		fuel:          8.3,
		distTravelled: 109.2,
		avgSpeed:      20.1,
	}

	c2 := Audi{
		fuel:          10.3,
		distTravelled: 111.38,
	}

	var t motorMilage

	t = BMW{10.10, 111.111, 111.11}

	fmt.Printf("milage of BMW: %f\n total run time of BMW: %f\n", t.milage(), t.totalTimeRun())

	fmt.Println(c1.milage(), c2.milage())
	fmt.Println("total time taken by BMW", c1.totalTimeRun())
}
