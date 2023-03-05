package main

import (
	"fmt"
)

type motorMilage interface {
	milage() float64
}

type BMW struct {
	fuel          float64
	avgSpeed      string
	distTravelled float64
}

type Audi struct {
	fuel          float64
	distTravelled float64
}

func (c BMW) milage() float64 {
	return c.distTravelled / c.fuel
}

func (c Audi) milage() float64 {
	return c.distTravelled / c.fuel
}

func main() {
	fmt.Println("hi")
	c1 := BMW{
		fuel:          8.3,
		distTravelled: 109.2,
		avgSpeed:      "20 kmph",
	}

	c2 := Audi{
		fuel:          10.3,
		distTravelled: 111.38,
	}

	fmt.Println(c1.milage(), c2.milage())

}
