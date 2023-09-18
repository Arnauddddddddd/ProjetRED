package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {

	for run && !(rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth/2 - buttonMenu.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + buttonMenu.Width/2 && screenHeight/2 - buttonMenu.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/2 + buttonMenu.Height/2)) {
		menu()
	}
	for run && 
	!(rl.IsMouseButtonDown(rl.MouseLeftButton) && (buttonPlay.Width < rl.GetMouseX() && rl.GetMouseX() < buttonPlay.Width*2 && screenHeight/4 - buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - buttonPlay.Height/2 +  buttonPlay.Height || screenWidth/2 - buttonPlay.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + buttonPlay.Width/2 && screenHeight/4 - buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - buttonPlay.Height/2 +  buttonPlay.Height || (screenWidth - buttonPlay.Width*2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth - buttonPlay.Width && screenHeight/4 - buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - buttonPlay.Height/2 +  buttonPlay.Height))){
		classSelector()
	}
	for run {
		input()
		update()
		render()
	}
	quit()
}
