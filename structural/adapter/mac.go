package main

import "fmt"

type Mac struct{}

func (m *Mac) InsertIntoLightingPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}
