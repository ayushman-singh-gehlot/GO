package main

import (
	"packages/gadget"
)

func playlist(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Samjhawaan", "Bekhayali", "Manjha"}
	playlist(player, mixtape)

	// 	player=gadget.TapeRecoder{}
	// 	playlist(player,mixtape)
	// this gives us error cause 1st argument has to be of type TapePlayer not TapeRecoder
	// we have solved this problem in fix_play_song.go

}
