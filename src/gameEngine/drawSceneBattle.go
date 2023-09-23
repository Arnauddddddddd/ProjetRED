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
	
	rl.DrawTexturePro(engine.monster[engine.monsterBattle].sprite, rl.NewRectangle(engine.monster[engine.monsterBattle].Src.X+float32(engine.monster[engine.monsterBattle].Xstart), engine.monster[engine.monsterBattle].Src.Y, engine.monster[engine.monsterBattle].Src.Width*-1, engine.monster[engine.monsterBattle].Src.Height), rl.NewRectangle(screenWidth-400-engine.monster[engine.monsterBattle].Dest.Width/2, screenHeight/2-50-engine.monster[engine.monsterBattle].Dest.Height/2, engine.monster[engine.monsterBattle].Dest.Width*6, engine.monster[engine.monsterBattle].Dest.Height*6), rl.NewVector2(engine.monster[engine.monsterBattle].Dest.Width, engine.monster[engine.monsterBattle].Dest.Height), 0, rl.White)

	engine.player.Src.X = 16*3
	rl.DrawTexturePro(engine.player.Sprite, engine.player.Src, rl.NewRectangle(350, 500, engine.player.Dest.Width*6, engine.player.Dest.Height*6), rl.NewVector2(engine.player.Dest.Width, engine.player.Dest.Height), 0, rl.White)

	rl.DrawTexturePro(engine.sprite.invBar, rl.NewRectangle(0, 0, float32(engine.sprite.invBar.Width), float32(engine.sprite.invBar.Height)), rl.NewRectangle(float32(screenWidth/2 - engine.sprite.invBar.Width/2 + 60), float32(screenHeight/2-engine.sprite.invBar.Height/2 + 465), 413, 216), rl.NewVector2(200, 200), 0, rl.White)

	rl.DrawTexturePro(engine.sprite.buttonBattle, rl.NewRectangle(345, 40, 380, 75), rl.NewRectangle( 280, screenHeight+ 45, 460, 110), rl.NewVector2(200, 200), 0, rl.White)
	rl.DrawTexturePro(engine.sprite.buttonBattle, rl.NewRectangle(345, 40+95, 380, 75), rl.NewRectangle( screenWidth-290, screenHeight+ 45, 460, 110), rl.NewVector2(200, 200), 0, rl.White)
	
	rl.DrawTextEx(engine.fontText, "Attack", rl.NewVector2( 215, 785), 45, 0, rl.Black)

	rl.DrawTextEx(engine.fontText, "Fatality", rl.NewVector2( screenWidth-360, 785), 45, 0, rl.Black)

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