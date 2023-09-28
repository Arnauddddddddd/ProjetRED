package gameEngine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func (engine *EngineStruct) menuFunc() {
	engine.run = !rl.WindowShouldClose()

	rl.BeginDrawing()

	rl.ClearBackground(engine.bgColor)

	engine.framCount++

	if engine.framCount % 10 == 1 { engine.menu.frameCount++ }
	if engine.menu.frameCount > 19 { engine.menu.frameCount = 0 }
	engine.menu.Src.X = engine.menu.Src.Width * float32(engine.menu.frameCount)// background changement de frame (voir dans classSelector)

	rl.DrawTexturePro(engine.menu.sprite, engine.menu.Src, engine.menu.Dest, rl.NewVector2(engine.menu.Dest.Width, engine.menu.Dest.Height), 0, rl.White) // background

	rl.DrawTextEx(engine.fontText, "Monster Slayer", rl.NewVector2(float32(screenWidth/2-306), 115), 80, 0, rl.White)
	rl.DrawTextEx(engine.fontText, "Monster Slayer", rl.NewVector2(float32(screenWidth/2-304), 115), 80, 0, rl.Black) // afficher le nom du jeu
	

    if screenWidth/2 - engine.sprite.buttonMenu.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + engine.sprite.buttonMenu.Width/2 && screenHeight/2 - engine.sprite.buttonMenu.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/2 + engine.sprite.buttonMenu.Height/2 { // si on passe la souris sur le bouton on change l'image du buttton pour un bouton pressé
		rl.DrawTexture(engine.sprite.buttonMenuPressed, screenWidth/2 - engine.sprite.buttonMenuPressed.Width/2, screenHeight/2 - engine.sprite.buttonMenu.Height/2, rl.White)
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			engine.menuPlay = false // si on appuie sur le bouton on sort du menu
		}
    } else {
        rl.DrawTexture(engine.sprite.buttonMenu, screenWidth/2 - engine.sprite.buttonMenu.Width/2, screenHeight/2 - engine.sprite.buttonMenu.Height/2, rl.White)
    }

	rl.EndDrawing()
}
