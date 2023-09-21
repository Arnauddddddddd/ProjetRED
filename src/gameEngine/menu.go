package gameEngine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func menu(engine *EngineStruct) {
	engine.run = !rl.WindowShouldClose()

	rl.BeginDrawing()

	rl.ClearBackground(engine.bgColor)
    if screenWidth/2 - engine.sprite.buttonMenu.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + engine.sprite.buttonMenu.Width/2 && screenHeight/2 - engine.sprite.buttonMenu.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/2 + engine.sprite.buttonMenu.Height/2 {
        rl.DrawTexture(engine.sprite.buttonMenuPressed, screenWidth/2 - engine.sprite.buttonMenuPressed.Width/2, screenHeight/2 - engine.sprite.buttonMenu.Height/2, rl.White)
    } else {
        rl.DrawTexture(engine.sprite.buttonMenu, screenWidth/2 - engine.sprite.buttonMenu.Width/2, screenHeight/2 - engine.sprite.buttonMenu.Height/2, rl.White)
    }

	rl.EndDrawing()
}
