package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	
	for engine.run && !(rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth/2 - sprite.buttonMenu.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + sprite.buttonMenu.Width/2 && screenHeight/2 - sprite.buttonMenu.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/2 + sprite.buttonMenu.Height/2)) {
		menu()
	}

	for engine.run && engine.menuSelector {
	classSelector()
	}

	for engine.run {
		input()
		update()
		render()
	}
	quit()
}
