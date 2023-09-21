package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func classSelector() {
	engine.run = !rl.WindowShouldClose()
	rl.BeginDrawing()

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
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Red Bull", "Potion", "A speed potion", rl.LoadTexture("../texture/PropsInPixels_16x/RedBull.png"), 0, 0, 1, true, false})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true})
        engine.character.inventory = append(engine.character.inventory, itemStruct{"Strengt item", "Item", "A strenghting potion", rl.LoadTexture("../texture/PropsInPixels_16x/star.png"), 100, 0, 0, false, true})
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