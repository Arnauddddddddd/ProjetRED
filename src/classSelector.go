package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func classSelector() {
	engine.run = !rl.WindowShouldClose()
	rl.BeginDrawing()

    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (sprite.buttonPlay.Width < rl.GetMouseX() && rl.GetMouseX() < sprite.buttonPlay.Width*2 && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 + sprite.buttonPlay.Height) {
        engine.menuSelector = false
    }
    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth/2 - sprite.buttonPlay.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + sprite.buttonPlay.Width/2 && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height) {
        engine.menuSelector = false
        character.name = "MOI"
        character.hp = 150
        character.damage = 50
        character.speed = 10
        character.hpMax = 200
        character.gold = 15
        character.alive = true
        character.inventory = append(character.inventory, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true})
        character.inventory = append(character.inventory, itemStruct{"Red Bull", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/RedBull.png"), 0, 0, 1, true, false})
        character.inventory = append(character.inventory, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/potion.png"), 0, 40, 0, true, true})
        character.inventory = append(character.inventory, itemStruct{"Heal Potion", "Potion", "A healing potion", rl.LoadTexture("../texture/PropsInPixels_16x/star.png"), 0, 40, 0, true, true})
    }
    if rl.IsMouseButtonDown(rl.MouseLeftButton) && (screenWidth - sprite.buttonPlay.Width*2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth - sprite.buttonPlay.Width && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height) {
        engine.menuSelector = false
        
    }
	if sprite.buttonPlay.Width < rl.GetMouseX() && rl.GetMouseX() < sprite.buttonPlay.Width*2 && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height{
        rl.DrawTexture(sprite.buttonPlay, sprite.buttonPlay.Width, screenHeight/4 - sprite.buttonPlay.Height/2, rl.White)
    } else {
        rl.DrawTexture(sprite.buttonPlayPressed, sprite.buttonPlay.Width, screenHeight/4 - sprite.buttonPlay.Height/2, rl.White)
    }

	if screenWidth/2 - sprite.buttonPlay.Width/2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth/2 + sprite.buttonPlay.Width/2 && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height{
        rl.DrawTexture(sprite.buttonPlay, screenWidth/2 - sprite.buttonPlay.Width/2, screenHeight/4 - sprite.buttonPlay.Height/2, rl.White)
    } else {
        rl.DrawTexture(sprite.buttonPlayPressed, screenWidth/2 - sprite.buttonPlay.Width/2, screenHeight/4 - sprite.buttonPlay.Height/2, rl.White)
    }

	if screenWidth - sprite.buttonPlay.Width*2 < rl.GetMouseX() && rl.GetMouseX() < screenWidth - sprite.buttonPlay.Width && screenHeight/4 - sprite.buttonPlay.Height/2 < rl.GetMouseY() && rl.GetMouseY() < screenHeight/4 - sprite.buttonPlay.Height/2 +  sprite.buttonPlay.Height{
        rl.DrawTexture(sprite.buttonPlay, screenWidth - sprite.buttonPlay.Width*2, screenHeight/4 - sprite.buttonPlay.Height/2, rl.White)
    } else {
        rl.DrawTexture(sprite.buttonPlayPressed, screenWidth - sprite.buttonPlay.Width*2, screenHeight/4 - sprite.buttonPlay.Height/2, rl.White)
    }

	rl.ClearBackground(engine.bgColor)
	rl.EndDrawing()
}