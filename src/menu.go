package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	
	"strconv"
)

func menu() {
	run = !rl.WindowShouldClose()

	rl.BeginDrawing()

	rl.ClearBackground(bgColor)
    if screenWidth/2 - buttonMenu.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + buttonMenu.Width/2 && screenHeight/2 - buttonMenu.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/2 + buttonMenu.Height/2 {
        rl.DrawTexture(buttonMenuPressed, screenWidth/2 - buttonMenuPressed.Width/2, screenHeight/2 - buttonMenu.Height/2, rl.White)
    } else {
        rl.DrawTexture(buttonMenu, screenWidth/2 - buttonMenu.Width/2, screenHeight/2 - buttonMenu.Height/2, rl.White)
    }
	
    rl.DrawText(strconv.Itoa(int(rl.GetMouseX())) + " " + strconv.Itoa(int(rl.GetMouseY())), 0, 0, 20, rl.Black)

	rl.EndDrawing()
}
