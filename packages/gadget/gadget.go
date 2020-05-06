package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("stopped !!")
}

type TapeRecoder struct {
	Batteries string
}

func (t TapeRecoder) Play(song string) {
	fmt.Println("playing", song)
}
func (t TapeRecoder) Stop() {
	fmt.Println("stopped !!")
}
func (t TapeRecoder) Record() {
	fmt.Println("recording")
}
