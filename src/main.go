package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	
	for engine.run && !(rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth/2 - sprite.buttonMenu.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + sprite.buttonMenu.Width/2 && screenHeight/2 - sprite.buttonMenu.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/2 + sprite.buttonMenu.Height/2)) {
		menu()
	}

	for engine.run && 
	!(rl.IsMouseButtonDown(rl.MouseLeftButton) && (sprite.buttonPlay.Width < rl.GetMouseX() && rl.GetMouseX() < sprite.buttonPlay.Width*2 && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height || screenWidth/2 - sprite.buttonPlay.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + sprite.buttonPlay.Width/2 && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height || (screenWidth - sprite.buttonPlay.Width*2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth - sprite.buttonPlay.Width && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height))){
		classSelector()
	}

	for engine.run {
		input()
		update()
		render()
	}
	quit()
}
