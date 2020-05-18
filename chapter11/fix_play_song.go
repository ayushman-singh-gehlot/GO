package main

import (
	"packages/gadget"
)

type instumentInterface interface {
	Play(string)
	Stop()
}

func playlist(device instumentInterface, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	var device instumentInterface
	mixtape := []string{"Samjhawaan", "Bekhayali", "Manjha"}

	device = gadget.TapePlayer{}
	playlist(device, mixtape)

	device = gadget.TapeRecoder{}
	playlist(device, mixtape)

}
