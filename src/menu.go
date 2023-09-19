package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func menu() {
	engine.run = !rl.WindowShouldClose()

	rl.BeginDrawing()

	rl.ClearBackground(engine.bgColor)
    if screenWidth/2 - sprite.buttonMenu.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + sprite.buttonMenu.Width/2 && screenHeight/2 - sprite.buttonMenu.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/2 + sprite.buttonMenu.Height/2 {
        rl.DrawTexture(sprite.buttonMenuPressed, screenWidth/2 - sprite.buttonMenuPressed.Width/2, screenHeight/2 - sprite.buttonMenu.Height/2, rl.White)
    } else {
        rl.DrawTexture(sprite.buttonMenu, screenWidth/2 - sprite.buttonMenu.Width/2, screenHeight/2 - sprite.buttonMenu.Height/2, rl.White)
    }

	rl.EndDrawing()
}
