package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)
/*
type personne struct {
	name string
	pv int
	speed int
}*/

func main() {

	for run && !(rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth/2 - buttonMenu.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + buttonMenu.Width/2 && screenHeight/2 - buttonMenu.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/2 + buttonMenu.Height/2)) {
		menu()
	}

	for run {
		input()
		update()
		render()
	}
	quit()
}
