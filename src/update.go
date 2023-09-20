package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func update() {
	engine.run = !rl.WindowShouldClose()
	if !character.showInventory {
		if !engine.battle {

			player.Src.Y = player.Src.Width * float32(player.Frame)

			for i := range bord.colisionList {
				if player.Up {
					nextBlock := []float32{player.Dest.X, player.Dest.Y - player.Src.Height}
					nextBlock2 := []float32{player.Dest.X - player.Src.Width, player.Dest.Y - player.Src.Height}
					if ((bord.colisionList[i][0]-16 < nextBlock[0]-2 && nextBlock[0]-2 < bord.colisionList[i][0]) && (bord.colisionList[i][1]-16 < nextBlock[1] && nextBlock[1] < bord.colisionList[i][1])) ||
					((bord.colisionList[i][0]-16 < nextBlock2[0]+2 && nextBlock2[0]+2 < bord.colisionList[i][0]) && (bord.colisionList[i][1]-16 < nextBlock2[1] && nextBlock2[1] < bord.colisionList[i][1])) {
						player.Moving = false
					}
				}
				if player.Down {
					nextBlock := []float32{player.Dest.X, player.Dest.Y}
					nextBlock2 := []float32{player.Dest.X - player.Src.Width, player.Dest.Y}
					if ((bord.colisionList[i][0]-16 < nextBlock[0]-2 && nextBlock[0]-2 < bord.colisionList[i][0]) && (bord.colisionList[i][1]-16 < nextBlock[1] && nextBlock[1] < bord.colisionList[i][1])) ||
					((bord.colisionList[i][0]-16 < nextBlock2[0]+2 && nextBlock2[0]+2 < bord.colisionList[i][0]) && (bord.colisionList[i][1]-16 < nextBlock2[1] && nextBlock2[1] < bord.colisionList[i][1])) {
						player.Moving = false
					}
				}
				if player.Left {
					nextBlock := []float32{player.Dest.X - player.Src.Width, player.Dest.Y - player.Src.Height}
					nextBlock2 := []float32{player.Dest.X - player.Src.Width, player.Dest.Y}
					if ((bord.colisionList[i][0]-16 < nextBlock[0] && nextBlock[0] < bord.colisionList[i][0]) && (bord.colisionList[i][1]-16 < nextBlock[1]+2 && nextBlock[1]+2 < bord.colisionList[i][1])) ||
					((bord.colisionList[i][0]-16 < nextBlock2[0] && nextBlock2[0] < bord.colisionList[i][0]) && (bord.colisionList[i][1]-16 < nextBlock2[1]-2 && nextBlock2[1]-2 < bord.colisionList[i][1])){
						player.Moving = false
					}
				}
				if player.Right {
					nextBlock := []float32{player.Dest.X, player.Dest.Y - player.Src.Height}
					nextBlock2 := []float32{player.Dest.X, player.Dest.Y}
					if ((bord.colisionList[i][0]-16 < nextBlock[0] && nextBlock[0] < bord.colisionList[i][0]) && (bord.colisionList[i][1]-16 < nextBlock[1]+2 && nextBlock[1]+2 < bord.colisionList[i][1])) ||
					((bord.colisionList[i][0]-16 < nextBlock2[0] && nextBlock2[0] < bord.colisionList[i][0]) && (bord.colisionList[i][1]-16 < nextBlock2[1]-2 && nextBlock2[1]-2 < bord.colisionList[i][1])) {
						player.Moving = false
					}
				}
			}
			if player.Moving {
				if player.Up { player.Dest.Y -= player.Speed }
				if player.Down { player.Dest.Y += player.Speed }
				if player.Left { player.Dest.X -= player.Speed }
				if player.Right { player.Dest.X += player.Speed }
				if engine.framCount % 8 == 1 { player.Frame++ }
			}

			engine.framCount++
			if player.Frame > 3 { player.Frame = 0 }
			if !player.Moving && player.Frame > 1 { player.Frame = 0 }

			player.Src.Y = player.Src.Width * float32(player.Frame)
			player.Src.X = player.Src.Height * float32(player.Dir)

			for i := 0; i < len(monster); i++ {
				if engine.framCount % 8 == 1 { monster[i].frameCount++ }
				if monster[i].frameCount > 3 { monster[i].frameCount = 0 }
				monster[i].Src.X = monster[i].Src.Width * float32(monster[i].frameCount)
				if rl.CheckCollisionRecs(player.Dest, monster[i].Dest) {
					engine.battle = true
					engine.monsterBattle = monster[i]
				}
			}

			rl.UpdateMusicStream(engine.music)
			if engine.musicPaused {
				rl.PauseMusicStream(engine.music)
			} else {
				rl.ResumeMusicStream(engine.music)
			}

			engine.cam.Target = rl.NewVector2(float32(player.Dest.X - (player.Dest.Width / 2)), float32(player.Dest.Y - (player.Dest.Height/2)))

			player.Moving = false
			player.Up, player.Down, player.Right, player.Left = false, false, false, false
		} else {
			if rl.IsKeyPressed(rl.KeyU) {
				engine.battle = false
			}
			//play := "player"
			/*
			if engine.monsterBattle.speed > character.speed {
				play = "monster"
			}

			if play == "player" {
				for !rl.IsMouseButtonDown(rl.MouseLeftButton) {
					time.Sleep(2 * time.Second)
				}
				engine.monsterBattle.hp -= character.damage
				if engine.monsterBattle.hp <= 0 {
					engine.monsterBattle.alive = false
					engine.battle = false
					character.hp += engine.monsterBattle.damage
				}
				character.hp -= engine.monsterBattle.damage
			} else {
				character.hp -= engine.monsterBattle.damage
				if character.hp <= 0 {
					character.alive = false
					engine.battle = false
					engine.monsterBattle.hp += character.damage
				}
				for !rl.IsMouseButtonDown(rl.MouseLeftButton) {
					continue
				}
				engine.monsterBattle.hp -= character.damage
			}*/
		}
	} else {
		inventorySelector()
	}
}
