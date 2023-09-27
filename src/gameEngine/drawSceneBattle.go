package gameEngine

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawSceneBattle(engine *EngineStruct) {

	rl.BeginDrawing()
	rl.ClearBackground(engine.bgColor)

	rl.DrawTexture(engine.sprite.bgForest, 0, 0, rl.White)

	rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*4, 0, 16*3, 16*2), rl.NewRectangle(210, 150, 16*3*3+20, 16*2*3), rl.NewVector2(16*3*3+20, 16*2*3), 0, rl.White)
	rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*4, 0, 16*3, 16*2), rl.NewRectangle(1550, 150, 16*3*3+20, 16*2*3), rl.NewVector2(16*3*3+20, 16*2*3), 0, rl.White)

	rl.DrawTextEx(engine.fontNum, strconv.Itoa(engine.character.hp), rl.NewVector2(112, 87), 26, 0, rl.Black)
	rl.DrawTextEx(engine.fontNum, strconv.Itoa(engine.monster[engine.battle.monsterBattle].hp), rl.NewVector2(1410, 87), 26, 0, rl.Black)

	monsterSrc := rl.NewRectangle(engine.monster[engine.battle.monsterBattle].Src.X+float32(engine.monster[engine.battle.monsterBattle].Xstart), engine.monster[engine.battle.monsterBattle].Src.Y, engine.monster[engine.battle.monsterBattle].Src.Width*-1, engine.monster[engine.battle.monsterBattle].Src.Height)
	monsterDest := rl.NewRectangle(screenWidth/2+300, screenHeight/2-engine.monster[engine.battle.monsterBattle].Src.Height/2, engine.monster[engine.battle.monsterBattle].Src.Width*float32(engine.monster[engine.battle.monsterBattle].increase), engine.monster[engine.battle.monsterBattle].Src.Height*float32(engine.monster[engine.battle.monsterBattle].increase))
	rl.DrawTexturePro(engine.monster[engine.battle.monsterBattle].sprite, monsterSrc, monsterDest, rl.NewVector2(engine.monster[engine.battle.monsterBattle].Dest.Width, engine.monster[engine.battle.monsterBattle].Dest.Height), 0, rl.White)

	engine.player.Src.X = 16*3
	rl.DrawTexturePro(engine.player.Sprite, engine.player.Src, rl.NewRectangle(350, 500, engine.player.Dest.Width*6, engine.player.Dest.Height*6), rl.NewVector2(engine.player.Dest.Width, engine.player.Dest.Height), 0, rl.White)

	if engine.player.showHud {
		rl.DrawTexturePro(engine.sprite.invBar, rl.NewRectangle(0, 0, float32(engine.sprite.invBar.Width), float32(engine.sprite.invBar.Height)), rl.NewRectangle(float32(screenWidth/2 - engine.sprite.invBar.Width/2 + 60), float32(screenHeight/2-engine.sprite.invBar.Height/2 + 465), 413, 216), rl.NewVector2(200, 200), 0, rl.White)

		rl.DrawTexturePro(engine.battle.buttonBattleAttack, rl.NewRectangle(345, 40, 380, 75), rl.NewRectangle( 280, screenHeight+ 45, 460, 110), rl.NewVector2(200, 200), 0, rl.White)
		rl.DrawTexturePro(engine.battle.buttonBattleFattality, rl.NewRectangle(345, 40+95, 380, 75), rl.NewRectangle( screenWidth-290, screenHeight+ 45, 460, 110), rl.NewVector2(200, 200), 0, rl.White)

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
	}

	if engine.battle.slash && !engine.playerTurn {
		rl.DrawTexturePro(engine.battle.slashSprite, engine.battle.slashSrc, rl.NewRectangle(480, 600, 70*5, 39*5), rl.NewVector2(70*4, 39*4), 0, rl.White)
	} else if engine.battle.slash && engine.playerTurn {
		rl.DrawTexturePro(engine.battle.slashSprite, engine.battle.slashSrc, rl.NewRectangle(1300, 600, 70*5, 39*5), rl.NewVector2(70*4, 39*4), 0, rl.White)
	}

	if engine.character.showText {
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*4, 0, 16, 16*2), rl.NewRectangle(310, 880, 48*4, 32*2*4), rl.NewVector2(48*4, 32*2*4), 0, rl.White)
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*5, 0, 16, 16*2), rl.NewRectangle(48*4+310, 880, 48*4, 32*2*4), rl.NewVector2(48*4, 32*2*4), 0, rl.White)
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*5, 0, 16, 16*2), rl.NewRectangle(48*4*2+310, 880, 48*4, 32*2*4), rl.NewVector2(48*4, 32*2*4), 0, rl.White)
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*5, 0, 16, 16*2), rl.NewRectangle(48*4*3+310, 880, 48*4, 32*2*4), rl.NewVector2(48*4, 32*2*4), 0, rl.White)
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*5, 0, 16, 16*2), rl.NewRectangle(48*4*4+310, 880, 48*4, 32*2*4), rl.NewVector2(48*4, 32*2*4), 0, rl.White)
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*5, 0, 16, 16*2), rl.NewRectangle(48*4*5+310, 880, 48*4, 32*2*4), rl.NewVector2(48*4, 32*2*4), 0, rl.White)
		rl.DrawTexturePro(engine.textBox.sprite, rl.NewRectangle(16*6, 0, 16, 16*2), rl.NewRectangle(48*4*6+310, 880, 48*4, 32*2*4), rl.NewVector2(48*4, 32*2*4), 0, rl.White)
		if engine.textBox.frameCountSpace % 2 == 0 {
			rl.DrawTexturePro(engine.textBox.space, rl.NewRectangle(12*4, 12*5, 12, 12), rl.NewRectangle(48*6+990, 765, 12*2*4, 12*2*4), rl.NewVector2(12, 12), 0, rl.White)
		} else {
			rl.DrawTexturePro(engine.textBox.space, rl.NewRectangle(12*5, 12*5, 12, 12), rl.NewRectangle(48*6+990, 765, 12*2*4, 12*2*4), rl.NewVector2(12, 12), 0, rl.White)
		}
		if engine.textBox.frameCountText == len(engine.textBox.textPrint) && len(engine.textBox.textToPrint) > len(engine.textBox.textPrint) {
			engine.textBox.textPrint += string(engine.textBox.textToPrint[engine.textBox.frameCountText])
		}
		if engine.textBox.textPrint == engine.textBox.textToPrint {
			engine.textBox.textWriting = false
		}
		if len(engine.textBox.textToPrint) > len(engine.textBox.textPrint) || !engine.textBox.textWriting {
			rl.DrawTextEx(engine.fontText, engine.textBox.textPrint, rl.NewVector2(230, 690), 50, 0, rl.Black)
		}
	}

	rl.EndDrawing()

}