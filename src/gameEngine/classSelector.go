package gameEngine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func classSelector(engine *EngineStruct) {
	engine.run = !rl.WindowShouldClose()
	rl.BeginDrawing()

    engine.framCount++

    if engine.framCount % 10 == 1 { engine.menu.frameCount++ }
	if engine.menu.frameCount > 19 { engine.menu.frameCount = 0 }
	engine.menu.Src.X = engine.menu.Src.Width * float32(engine.menu.frameCount)

	rl.DrawTexturePro(engine.menu.sprite, engine.menu.Src, engine.menu.Dest, rl.NewVector2(engine.menu.Dest.Width, engine.menu.Dest.Height), 0, rl.White)

    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (engine.sprite.buttonPlay.Width < rl.GetMouseX() && rl.GetMouseX() < engine.sprite.buttonPlay.Width*2 && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 + engine.sprite.buttonPlay.Height) {
        engine.menuSelector = false
    }
    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth/2 - engine.sprite.buttonPlay.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + engine.sprite.buttonPlay.Width/2 && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 +  engine.sprite.buttonPlay.Height) {
        engine.menuSelector = false
        engine.character.name = "MOI"
        engine.character.hp = 150
        engine.character.damage = 50
        engine.character.damageBase = 50
        engine.character.speed = 10
        engine.character.hpMax = 200
        engine.character.gold = 15
        engine.character.alive = true
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true, 10})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Red Bull", "Potion", "A speed potion", rl.LoadTexture("../texture/PropsInPixels_16x/RedBull.png"), 0, 0, 3, true, false, 10})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true, 10})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true, 10})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Red Bull", "Potion", "A speed potion", rl.LoadTexture("../texture/PropsInPixels_16x/RedBull.png"), 0, 0, 3, true, false, 10})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true, 10})
    }
    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth - engine.sprite.buttonPlay.Width*2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth - engine.sprite.buttonPlay.Width && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 +  engine.sprite.buttonPlay.Height) {
        engine.menuSelector = false
        
    }
	if engine.sprite.buttonPlay.Width < rl.GetMouseX() && rl.GetMouseX() < engine.sprite.buttonPlay.Width*2 && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 +  engine.sprite.buttonPlay.Height{
        rl.DrawTexture(engine.sprite.buttonPlay, engine.sprite.buttonPlay.Width, screenHeight/4 - engine.sprite.buttonPlay.Height/2, rl.White)
    } else {
        rl.DrawTexture(engine.sprite.buttonPlayPressed, engine.sprite.buttonPlay.Width, screenHeight/4 - engine.sprite.buttonPlay.Height/2, rl.White)
    }

	if screenWidth/2 - engine.sprite.buttonPlay.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + engine.sprite.buttonPlay.Width/2 && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 +  engine.sprite.buttonPlay.Height{
        rl.DrawTexture(engine.sprite.buttonPlay, screenWidth/2 - engine.sprite.buttonPlay.Width/2, screenHeight/4 - engine.sprite.buttonPlay.Height/2, rl.White)
    } else {
        rl.DrawTexture(engine.sprite.buttonPlayPressed, screenWidth/2 - engine.sprite.buttonPlay.Width/2, screenHeight/4 - engine.sprite.buttonPlay.Height/2, rl.White)
    }

	if screenWidth - engine.sprite.buttonPlay.Width*2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth - engine.sprite.buttonPlay.Width && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 +  engine.sprite.buttonPlay.Height{
        rl.DrawTexture(engine.sprite.buttonPlay, screenWidth - engine.sprite.buttonPlay.Width*2, screenHeight/4 - engine.sprite.buttonPlay.Height/2, rl.White)
    } else {
        rl.DrawTexture(engine.sprite.buttonPlayPressed, screenWidth - engine.sprite.buttonPlay.Width*2, screenHeight/4 - engine.sprite.buttonPlay.Height/2, rl.White)
    }

	rl.ClearBackground(engine.bgColor)
	rl.EndDrawing()
}