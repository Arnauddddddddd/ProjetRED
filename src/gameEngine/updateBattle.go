package gameEngine

import (

	rl "github.com/gen2brain/raylib-go/raylib"

	"fmt"
)

func updateBattle(engine *EngineStruct) {

	engine.run = !rl.WindowShouldClose()

	if !engine.character.showText {

		engine.framCount++

		if engine.framCount%8 == 1 {
			engine.monster[engine.battle.monsterBattle].frameCount++
		}
		if engine.monster[engine.battle.monsterBattle].frameCount > 3 {
			engine.monster[engine.battle.monsterBattle].frameCount = 0
		}
		engine.monster[engine.battle.monsterBattle].Src.X = engine.monster[engine.battle.monsterBattle].Src.Width * float32(engine.monster[engine.battle.monsterBattle].frameCount)

		if engine.framCount%2 == 1 && engine.battle.slash {
			engine.battle.slashFrameCount++
		}
		engine.battle.slashSrc.X = engine.battle.slashSrc.Width * float32(engine.battle.slashFrameCount)

		if engine.monster[engine.battle.monsterBattle].speed > engine.character.speed {
			engine.playerTurn = false
			engine.monster[engine.battle.monsterBattle].speed = 0
			engine.character.showText = true
			engine.textBox.textToPrint = "Attention il attaque il est plus rapide .. "
			engine.textBox.textPrint = ""
			engine.textBox.frameCountText = 0
			engine.textBox.textWriting = true
		} else {
			if engine.playerTurn {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) && (126 < rl.GetMouseX() && rl.GetMouseX() < 468 && 758 < rl.GetMouseY() && rl.GetMouseY() < 840) && engine.player.showHud{
					engine.player.showHud = false
					drawSceneBattle(engine)

					engine.character.showText = true
					engine.textBox.textToPrint = "Vous attaquez .. "
					engine.textBox.textPrint = ""
					engine.textBox.frameCountText = 0
					engine.textBox.textWriting = true

				}
				if !engine.character.showText && !engine.player.showHud && !engine.battle.slash {
					engine.battle.slashFrameCount = 0
					engine.battle.slash = true
				}
				if 126 < rl.GetMouseX() && rl.GetMouseX() < 468 && 758 < rl.GetMouseY() && rl.GetMouseY() < 840 {
					engine.battle.buttonBattleAttack = engine.battle.buttonBattlePressed[1]
				} else {
					engine.battle.buttonBattleAttack = engine.battle.buttonBattlePressed[0]
				}
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) && (1140 < rl.GetMouseX() && rl.GetMouseX() < 1492 && 758 < rl.GetMouseY() && rl.GetMouseY() < 840) {
					fmt.Println(rl.GetMouseX())
				}
				if 1140 < rl.GetMouseX() && rl.GetMouseX() < 1492 && 758 < rl.GetMouseY() && rl.GetMouseY() < 840 {
					engine.battle.buttonBattleFattality = engine.battle.buttonBattlePressed[1]
				} else {
					engine.battle.buttonBattleFattality = engine.battle.buttonBattlePressed[0]
				}
				if engine.battle.slashFrameCount > 17 && engine.battle.slash {
					engine.battle.slash = false
					engine.battle.slashFrameCount = 0
					engine.playerTurn = false
					engine.monster[engine.battle.monsterBattle].hp -= engine.character.damage
					drawSceneBattle(engine)
					engine.character.showText = true
					engine.textBox.textToPrint = "Attention il attaque .. "
					engine.textBox.textPrint = ""
					engine.textBox.frameCountText = 0
					engine.textBox.textWriting = true
				}
				if engine.monster[engine.battle.monsterBattle].hp <= 0 {
					engine.character.gold += engine.monster[engine.battle.monsterBattle].goldLoot
					engine.character.damage = engine.character.damageBase
					engine.monster[engine.battle.monsterBattle].alive = false
					engine.battle.inBattle = false
					engine.character.showText = false
				}
				if engine.character.hp <= 0 {
					engine.character.damage = engine.character.damageBase
					engine.character.alive = false
					engine.battle.inBattle = false
					engine.character.showText = false
				}
			} else {
				engine.battle.slash = true
				if engine.battle.slashFrameCount > 17 {
					engine.battle.slashFrameCount = 0
					engine.battle.slash = false
					engine.playerTurn = true
					engine.character.hp -= engine.monster[engine.battle.monsterBattle].damage
					drawSceneBattle(engine)
					engine.player.showHud = true
				}
			}
		}
	} else {
		engine.framCount++

		if engine.framCount%8 == 1 {
			engine.monster[engine.battle.monsterBattle].frameCount++
		}
		if engine.monster[engine.battle.monsterBattle].frameCount > 3 {
			engine.monster[engine.battle.monsterBattle].frameCount = 0
		}
		engine.monster[engine.battle.monsterBattle].Src.X = engine.monster[engine.battle.monsterBattle].Src.Width * float32(engine.monster[engine.battle.monsterBattle].frameCount)
		text(engine)
	}

}
