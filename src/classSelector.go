package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func classSelector() {
	run = !rl.WindowShouldClose()
	rl.BeginDrawing()

	if buttonPlay.Width < rl.GetMouseX() && rl.GetMouseX() < buttonPlay.Width*2 && screenHeight/4 - buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - buttonPlay.Height/2 +  buttonPlay.Height{
        rl.DrawTexture(buttonPlayPressed, buttonPlay.Width, screenHeight/4 - buttonPlay.Height/2, rl.White)
    } else {
        rl.DrawTexture(buttonPlay, buttonPlay.Width, screenHeight/4 - buttonPlay.Height/2, rl.White)
    }

	if screenWidth/2 - buttonPlay.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + buttonPlay.Width/2 && screenHeight/4 - buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - buttonPlay.Height/2 +  buttonPlay.Height{
        rl.DrawTexture(buttonPlayPressed, screenWidth/2 - buttonPlay.Width/2, screenHeight/4 - buttonPlay.Height/2, rl.White)
    } else {
        rl.DrawTexture(buttonPlay, screenWidth/2 - buttonPlay.Width/2, screenHeight/4 - buttonPlay.Height/2, rl.White)
    }

	if screenWidth - buttonPlay.Width*2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth - buttonPlay.Width && screenHeight/4 - buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - buttonPlay.Height/2 +  buttonPlay.Height{
        rl.DrawTexture(buttonPlayPressed, screenWidth - buttonPlay.Width*2, screenHeight/4 - buttonPlay.Height/2, rl.White)
    } else {
        rl.DrawTexture(buttonPlay, screenWidth - buttonPlay.Width*2, screenHeight/4 - buttonPlay.Height/2, rl.White)
    }

	rl.ClearBackground(bgColor)
	rl.EndDrawing()
}