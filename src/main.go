package main

import "mars-rover-navigation/src/console"

func main() {
	c, err := console.Wire()
	if err != nil {
		panic(err)
	}
	c.Start()
}
