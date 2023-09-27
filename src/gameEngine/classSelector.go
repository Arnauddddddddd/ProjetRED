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

    rl.DrawTextEx(engine.fontText, "Thief", rl.NewVector2(float32(engine.sprite.buttonPlay.Width-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+100)), 60, 0, rl.Black)
    rl.DrawTextEx(engine.fontText, "Slayer", rl.NewVector2(float32(screenWidth/2 - engine.sprite.buttonPlay.Width/2-30), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+100)), 60, 0, rl.Black)
    rl.DrawTextEx(engine.fontText, "Tank", rl.NewVector2(float32(screenWidth - engine.sprite.buttonPlay.Width*2-20), float32(screenHeight/4 - engine.sprite.buttonPlay.Height/2+100)), 60, 0, rl.Black)

    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (engine.sprite.buttonPlay.Width < rl.GetMouseX() && rl.GetMouseX() < engine.sprite.buttonPlay.Width*2 && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 + engine.sprite.buttonPlay.Height) {
        engine.menuSelector = false
        engine.character.name = "Thief"
        engine.character.hp = 150
        engine.character.damage = 50
        engine.character.damageBase = 50
        engine.character.speed = 10
        engine.character.hpMax = 150
        engine.character.gold = 15
        engine.character.alive = true
    }
    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth/2 - engine.sprite.buttonPlay.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + engine.sprite.buttonPlay.Width/2 && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 +  engine.sprite.buttonPlay.Height) {
        engine.menuSelector = false
        engine.character.name = "Slayer"
        engine.character.hp = 150
        engine.character.damage = 90
        engine.character.damageBase = 90
        engine.character.speed = 10
        engine.character.hpMax = 150
        engine.character.gold = 15
        engine.character.alive = true
        engine.character.inventory = append( engine.character.inventory, itemStruct{"Strengt item", "Fork", "By using this fork, you fight \nthe rest of the fight with it and will \ninflict an additional 50 damage on you.", rl.LoadTexture("../texture/shop/fork.png"), 0, 0, 0, false, true, false,1})
        engine.character.inventory = append( engine.character.inventory, itemStruct{"Strengt item", "Fork", "By using this fork, you fight \nthe rest of the fight with it and will \ninflict an additional 50 damage on you.", rl.LoadTexture("../texture/shop/fork.png"), 0, 0, 0, false, true, false,1})
    }
    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth - engine.sprite.buttonPlay.Width*2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth - engine.sprite.buttonPlay.Width && screenHeight/4 - engine.sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - engine.sprite.buttonPlay.Height/2 +  engine.sprite.buttonPlay.Height) {
        engine.menuSelector = false
        engine.character.name = "Tank"
        engine.character.hp = 250
        engine.character.damage = 50
        engine.character.damageBase = 50
        engine.character.speed = 10
        engine.character.hpMax = 250
        engine.character.gold = 15
        engine.character.alive = true
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Red Bull", "Potion", "A speed potion", rl.LoadTexture("../texture/shop/RedBull.png"), 0, 0, 3, true, false, true, 10})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Red Bull", "Potion", "A speed potion", rl.LoadTexture("../texture/shop/RedBull.png"), 0, 0, 3, true, false, true, 10})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Red Bull", "Potion", "A speed potion", rl.LoadTexture("../texture/shop/RedBull.png"), 0, 0, 3, true, false, true, 10})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Red Bull", "Potion", "A speed potion", rl.LoadTexture("../texture/shop/RedBull.png"), 0, 0, 3, true, false, true, 10})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Red Bull", "Potion", "A speed potion", rl.LoadTexture("../texture/shop/RedBull.png"), 0, 0, 3, true, false, true, 10})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Red Bull", "Potion", "A speed potion", rl.LoadTexture("../texture/shop/RedBull.png"), 0, 0, 3, true, false, true, 10})
   
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