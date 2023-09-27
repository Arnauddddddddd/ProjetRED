package gameEngine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func update(engine *EngineStruct) {
	engine.run = !rl.WindowShouldClose()
	if !engine.character.showText {
		if !engine.character.showInventory {

			if !engine.character.alive {
				engine.character.gold /= 2
				engine.character.hp = engine.character.hpMax
				engine.character.alive = true
				engine.player.Dest = rl.NewRectangle(1057, 1633, 16, 16)
			}

			engine.player.Src.Y = engine.player.Src.Width * float32(engine.player.Frame)

			for i := range engine.bord.colisionList {
				if engine.player.Up {
					nextBlock := []float32{engine.player.Dest.X, engine.player.Dest.Y - engine.player.Src.Height-5}
					nextBlock2 := []float32{engine.player.Dest.X - engine.player.Src.Width, engine.player.Dest.Y - engine.player.Src.Height-5}
					if ((engine.bord.colisionList[i][0]-16 < nextBlock[0]-2 && nextBlock[0]-2 < engine.bord.colisionList[i][0]) && (engine.bord.colisionList[i][1]-16 < nextBlock[1] && nextBlock[1] < engine.bord.colisionList[i][1])) ||
					((engine.bord.colisionList[i][0]-16 < nextBlock2[0]+2 && nextBlock2[0]+2 < engine.bord.colisionList[i][0]) && (engine.bord.colisionList[i][1]-16 < nextBlock2[1] && nextBlock2[1] < engine.bord.colisionList[i][1])) {
						engine.player.Moving = false
					}
				}
				if engine.player.Down {
					nextBlock := []float32{engine.player.Dest.X, engine.player.Dest.Y+5}
					nextBlock2 := []float32{engine.player.Dest.X - engine.player.Src.Width, engine.player.Dest.Y+5}
					if ((engine.bord.colisionList[i][0]-16 < nextBlock[0]-2 && nextBlock[0]-2 < engine.bord.colisionList[i][0]) && (engine.bord.colisionList[i][1]-16 < nextBlock[1] && nextBlock[1] < engine.bord.colisionList[i][1])) ||
					((engine.bord.colisionList[i][0]-16 < nextBlock2[0]+2 && nextBlock2[0]+2 < engine.bord.colisionList[i][0]) && (engine.bord.colisionList[i][1]-16 < nextBlock2[1] && nextBlock2[1] < engine.bord.colisionList[i][1])) {
						engine.player.Moving = false
					}
				}
				if engine.player.Left {
					nextBlock := []float32{engine.player.Dest.X - engine.player.Src.Width-5, engine.player.Dest.Y - engine.player.Src.Height}
					nextBlock2 := []float32{engine.player.Dest.X - engine.player.Src.Width-5, engine.player.Dest.Y}
					if ((engine.bord.colisionList[i][0]-16 < nextBlock[0] && nextBlock[0] < engine.bord.colisionList[i][0]) && (engine.bord.colisionList[i][1]-16 < nextBlock[1]+2 && nextBlock[1]+2 < engine.bord.colisionList[i][1])) ||
					((engine.bord.colisionList[i][0]-16 < nextBlock2[0] && nextBlock2[0] < engine.bord.colisionList[i][0]) && (engine.bord.colisionList[i][1]-16 < nextBlock2[1]-2 && nextBlock2[1]-2 < engine.bord.colisionList[i][1])){
						engine.player.Moving = false
					}
				}
				if engine.player.Right {
					nextBlock := []float32{engine.player.Dest.X+5, engine.player.Dest.Y - engine.player.Src.Height}
					nextBlock2 := []float32{engine.player.Dest.X+5, engine.player.Dest.Y}
					if ((engine.bord.colisionList[i][0]-16 < nextBlock[0] && nextBlock[0] < engine.bord.colisionList[i][0]) && (engine.bord.colisionList[i][1]-16 < nextBlock[1]+2 && nextBlock[1]+2 < engine.bord.colisionList[i][1])) ||
					((engine.bord.colisionList[i][0]-16 < nextBlock2[0] && nextBlock2[0] < engine.bord.colisionList[i][0]) && (engine.bord.colisionList[i][1]-16 < nextBlock2[1]-2 && nextBlock2[1]-2 < engine.bord.colisionList[i][1])) {
						engine.player.Moving = false
					}
				}
			}

			if engine.player.Moving {
				if engine.player.Up { engine.player.Dest.Y -= engine.player.Speed }
				if engine.player.Down { engine.player.Dest.Y += engine.player.Speed }
				if engine.player.Left { engine.player.Dest.X -= engine.player.Speed }
				if engine.player.Right { engine.player.Dest.X += engine.player.Speed }
				if engine.framCount % 8 == 1 { engine.player.Frame++ }
			}

			engine.framCount++
			if engine.player.Frame > 3 { engine.player.Frame = 0 }

			if !engine.player.Moving && engine.player.Frame > 1 { engine.player.Frame = 0 }

			engine.player.Src.Y = engine.player.Src.Width * float32(engine.player.Frame)
			engine.player.Src.X = engine.player.Src.Height * float32(engine.player.Dir)

			for i := 0; i < len(engine.monster); i++ {
				if engine.framCount % engine.monster[i].speedFrame == 1 { engine.monster[i].frameCount++ }
				if engine.monster[i].frameCount > engine.monster[i].frameNumber { engine.monster[i].frameCount = 0 }
				engine.monster[i].Src.X = engine.monster[i].Src.Width * float32(engine.monster[i].frameCount)
				if !engine.monster[i].alive {
					engine.monster[i].deadTime++
					if engine.monster[i].deadTime >= 1000 {
						engine.monster[i].alive = true
						engine.monster[i].hp = engine.monster[i].hpMax
						engine.monster[i].deadTime = 0
					}
				}
				if i < len(engine.monster)-1 {
					if (rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(engine.monster[i].Dest.X, engine.monster[i].Dest.Y - engine.monster[i].Dest.Height, 32, 32))) ||
					rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(engine.monster[i].Dest.X - engine.monster[i].Dest.Width, engine.monster[i].Dest.Y - engine.monster[i].Dest.Height, 32, 32))) ||
					rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(engine.monster[i].Dest.X, engine.monster[i].Dest.Y, 32, 32))) ||
					rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(engine.monster[i].Dest.X - engine.monster[i].Dest.Width, engine.monster[i].Dest.Y, 32, 32)))) && engine.monster[i].alive {
						//engine.battle.inBattle = true
						//engine.battle.monsterBattle = i
						//engine.player.showHud = true
						//engine.playerTurn = true
						// a remettre pour lancer les combats
					}
				}
				if i == len(engine.monster)-1 && engine.monster[len(engine.monster)-1].alive {
					if rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(1330, 100, 200, 400))) {
						engine.battle.inBattle = true
						engine.battle.monsterBattle = i
						engine.player.showHud = true
						engine.playerTurn = true
						// a remettre pour lancer les combats
					}
				}
			}

			if ((rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(engine.shopKeeper.Dest.X, engine.shopKeeper.Dest.Y - 32 - engine.shopKeeper.Dest.Height, 64, 64))) ||
				rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(engine.shopKeeper.Dest.X-32 - engine.shopKeeper.Dest.Width, engine.shopKeeper.Dest.Y - 32 - engine.shopKeeper.Dest.Height, 64, 64))) ||
				rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(engine.shopKeeper.Dest.X, engine.shopKeeper.Dest.Y, 64, 64))) ||
				rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(engine.shopKeeper.Dest.X-32 - engine.shopKeeper.Dest.Width, engine.shopKeeper.Dest.Y, 64, 64))))) ||
				((rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(686, 1153 - 32 - engine.shopKeeper.Dest.Height, 64, 64))) ||
				rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(686- 32 - engine.shopKeeper.Dest.Width, 1153 - 32 - engine.shopKeeper.Dest.Height, 64, 64))) ||
				rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(686, 1153, 64, 64))) ||
				rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(686 - engine.shopKeeper.Dest.Width, 1153, 64, 64)))) ||
				((rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(594, 780 - engine.shopKeeper.Dest.Height/2, 64, 64))) ||
				rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(594 - 32 - engine.shopKeeper.Dest.Width, 780 - 32 - engine.shopKeeper.Dest.Height, 64, 64))) ||
				rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(594, 780, 64, 64))) ||
				rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(594 - engine.shopKeeper.Dest.Width, 780, 64, 64))))) ||
				((rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(1030, 340 - 32 - engine.shopKeeper.Dest.Height, 64, 64))) ||
				rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(1030 - 32 - engine.shopKeeper.Dest.Width, 340 - 32 - engine.shopKeeper.Dest.Height, 64, 64))) ||
				rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(1030, 340, 64, 64))) ||
				rl.CheckCollisionRecs(engine.player.Dest, (rl.NewRectangle(1030 - engine.shopKeeper.Dest.Width, 340, 64, 64)))))) {
				engine.shop = true
			} else {
				engine.shop = false
			}
			//686, 1153
			if !engine.doorOpenKey {
				if 1020 < engine.player.Dest.X && engine.player.Dest.X < 1130 && 1510 < engine.player.Dest.Y && engine.player.Dest.Y < 1590 {
					for i := 0; i < len(engine.character.inventory); i++ {
						if engine.character.inventory[i].name == "Key" {
							engine.doorOpenKey = true
							engine.character.inventory = append(engine.character.inventory[:i], engine.character.inventory[i+1:]...)
						}
					}
				}
			} else {
				if 1020 < engine.player.Dest.X && engine.player.Dest.X < 1130 && 1460 < engine.player.Dest.Y && engine.player.Dest.Y < 1580 {
					engine.doorOpen = true
				} else {
					engine.doorOpen = false
				}
			}
			

			engine.cam.Target = rl.NewVector2(float32(engine.player.Dest.X - (engine.player.Dest.Width / 2)), float32(engine.player.Dest.Y - (engine.player.Dest.Height/2)))

			engine.player.Moving = false
			engine.player.Up, engine.player.Down, engine.player.Right, engine.player.Left = false, false, false, false
		} else {
			if engine.shop {
				shopSelector(engine)
			} else {
				inventorySelector(engine, 620, 365, 445, 455, 530 , 90)
			}
		}
	} else {
		text(engine)
	}
	rl.UpdateMusicStream(engine.music)
	if engine.musicPaused {
		rl.PauseMusicStream(engine.music)
	} else {
		rl.ResumeMusicStream(engine.music)
	}
}
