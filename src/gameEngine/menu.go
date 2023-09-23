package gameEngine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func menu(engine *EngineStruct) {
	engine.run = !rl.WindowShouldClose()

	rl.BeginDrawing()

	rl.ClearBackground(engine.bgColor)

	engine.framCount++

	if engine.framCount % 10 == 1 { engine.menu.frameCount++ }
	if engine.menu.frameCount > 19 { engine.menu.frameCount = 0 }
	engine.menu.Src.X = engine.menu.Src.Width * float32(engine.menu.frameCount)

	rl.DrawTexturePro(engine.menu.sprite, engine.menu.Src, engine.menu.Dest, rl.NewVector2(engine.menu.Dest.Width, engine.menu.Dest.Height), 0, rl.White)

	rl.DrawTextEx(engine.fontText, "Monster Slayer", rl.NewVector2(float32(screenWidth/2-306), 115), 80, 0, rl.White)
	rl.DrawTextEx(engine.fontText, "Monster Slayer", rl.NewVector2(float32(screenWidth/2-304), 115), 80, 0, rl.Black)
	

    if screenWidth/2 - engine.sprite.buttonMenu.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + engine.sprite.buttonMenu.Width/2 && screenHeight/2 - engine.sprite.buttonMenu.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/2 + engine.sprite.buttonMenu.Height/2 {
        rl.DrawTexture(engine.sprite.buttonMenuPressed, screenWidth/2 - engine.sprite.buttonMenuPressed.Width/2, screenHeight/2 - engine.sprite.buttonMenu.Height/2, rl.White)
    } else {
        rl.DrawTexture(engine.sprite.buttonMenu, screenWidth/2 - engine.sprite.buttonMenu.Width/2, screenHeight/2 - engine.sprite.buttonMenu.Height/2, rl.White)
    }

	rl.EndDrawing()
}
