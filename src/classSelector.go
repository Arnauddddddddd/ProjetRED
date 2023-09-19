package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

func classSelector() {
	engine.run = !rl.WindowShouldClose()
	rl.BeginDrawing()

    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (sprite.buttonPlay.Width < rl.GetMouseX() && rl.GetMouseX() < sprite.buttonPlay.Width*2 && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 + sprite.buttonPlay.Height) {
        fmt.Println("A")
        engine.menuSelector = false
    }
    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth/2 - sprite.buttonPlay.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + sprite.buttonPlay.Width/2 && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height) {
        fmt.Println("B")
        engine.menuSelector = false
    }
    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth - sprite.buttonPlay.Width*2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth - sprite.buttonPlay.Width && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height) {
        engine.menuSelector = false
        
    }
	if sprite.buttonPlay.Width < rl.GetMouseX() && rl.GetMouseX() < sprite.buttonPlay.Width*2 && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height{
        rl.DrawTexture(sprite.buttonPlayPressed, sprite.buttonPlay.Width, screenHeight/4 - sprite.buttonPlay.Height/2, rl.White)
    } else {
        rl.DrawTexture(sprite.buttonPlay, sprite.buttonPlay.Width, screenHeight/4 - sprite.buttonPlay.Height/2, rl.White)
    }

	if screenWidth/2 - sprite.buttonPlay.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + sprite.buttonPlay.Width/2 && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height{
        rl.DrawTexture(sprite.buttonPlayPressed, screenWidth/2 - sprite.buttonPlay.Width/2, screenHeight/4 - sprite.buttonPlay.Height/2, rl.White)
    } else {
        rl.DrawTexture(sprite.buttonPlay, screenWidth/2 - sprite.buttonPlay.Width/2, screenHeight/4 - sprite.buttonPlay.Height/2, rl.White)
    }

	if screenWidth - sprite.buttonPlay.Width*2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth - sprite.buttonPlay.Width && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height{
        rl.DrawTexture(sprite.buttonPlayPressed, screenWidth - sprite.buttonPlay.Width*2, screenHeight/4 - sprite.buttonPlay.Height/2, rl.White)
    } else {
        rl.DrawTexture(sprite.buttonPlay, screenWidth - sprite.buttonPlay.Width*2, screenHeight/4 - sprite.buttonPlay.Height/2, rl.White)
    }

	rl.ClearBackground(engine.bgColor)
	rl.EndDrawing()
}