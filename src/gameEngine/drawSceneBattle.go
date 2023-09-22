package gameEngine

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawSceneBattle(engine *EngineStruct) {

	rl.BeginDrawing()
	rl.ClearBackground(engine.bgColor)
	//rl.BeginMode2D(engine.cam)

	

	rl.DrawTexture(engine.sprite.bgForest, 0, 0, rl.White)
	rl.DrawText(strconv.Itoa(engine.character.hp), 50, 50, 30, rl.Black)
	rl.DrawText(strconv.Itoa(engine.monster[engine.monsterBattle].hp), int32(engine.player.Dest.X)+600, int32(engine.player.Dest.Y)-400, 30, rl.Black)
	rl.DrawTexturePro(engine.monster[engine.monsterBattle].sprite, engine.monster[engine.monsterBattle].Src, engine.monster[engine.monsterBattle].Dest, rl.NewVector2(engine.monster[engine.monsterBattle].Dest.Width, engine.monster[engine.monsterBattle].Dest.Height), 0, rl.White)

	//rl.DrawTexturePro(engine.sprite.buttonAttack, rl.NewRectangle(32, 32*6, 48, 32), rl.NewRectangle(engine.player.Dest.X-520, engine.player.Dest.Y+490, 230, 120), rl.NewVector2(200, 200), 0, rl.White)
	//rl.DrawTexturePro(engine.sprite.buttonAttack, rl.NewRectangle(32+48, 32*6, 48, 32), rl.NewRectangle(engine.player.Dest.X-520, engine.player.Dest.Y+490, 230, 120), rl.NewVector2(200, 200), 0, rl.White)
	//rl.DrawTexturePro(engine.sprite.buttonAttack, rl.NewRectangle(32*2+48*2, 32*3, 48, 32), rl.NewRectangle(engine.player.Dest.X+680, engine.player.Dest.Y+490, 230, 120), rl.NewVector2(200, 200), 0, rl.White)
	//rl.DrawTexturePro(engine.sprite.buttonAttack, rl.NewRectangle(32*2+48*3, 32*3, 48, 32), rl.NewRectangle(engine.player.Dest.X+680, engine.player.Dest.Y+490, 230, 120), rl.NewVector2(200, 200), 0, rl.White)
	
	rl.DrawTexturePro(engine.sprite.invBar, rl.NewRectangle(0, 0, float32(engine.sprite.invBar.Width), float32(engine.sprite.invBar.Height)), rl.NewRectangle(float32(screenWidth/2 - engine.sprite.invBar.Width/2 + 60), float32(screenHeight/2-engine.sprite.invBar.Height/2 + 465), 413, 216), rl.NewVector2(200, 200), 0, rl.White)

	rl.DrawTexturePro(engine.sprite.buttonBattle, rl.NewRectangle(345, 40, 380, 75), rl.NewRectangle( 280, screenHeight+ 45, 460, 110), rl.NewVector2(200, 200), 0, rl.White)
	rl.DrawTexturePro(engine.sprite.buttonBattle, rl.NewRectangle(345, 40+95, 380, 75), rl.NewRectangle( screenWidth-290, screenHeight+ 45, 460, 110), rl.NewVector2(200, 200), 0, rl.White)
	

	t := 0
	slot2 := 0
	for i := 0; i < len(engine.character.inventory); i++ {
		if i == 4 {
			slot2 = 1
			t = 0
		}
		rl.DrawTexturePro(engine.character.inventory[i].sprite, rl.NewRectangle(0, 0, float32(engine.character.inventory[i].sprite.Width), float32(engine.character.inventory[i].sprite.Height)), rl.NewRectangle(float32(screenWidth/2 - engine.sprite.invBar.Width/2 + 60 + 40 + int32((70+23)*t)), float32(screenHeight/2-engine.sprite.invBar.Height/2 + 465 + 38 + int32((64+23)*slot2)), 61, 61), rl.NewVector2(200, 200), 0, rl.White)
		t++
	}
	

	//rl.EndMode2D()
	rl.EndDrawing()

}