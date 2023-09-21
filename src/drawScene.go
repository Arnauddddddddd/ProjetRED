package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawScene() {
	if !engine.battle {
		engine.bord.colisionList = [][]float32{}
		for i := 0; i < len(engine.bord.tileMap); i++ {
			if engine.bord.tileMap[i] != 0 {
				engine.bord.tileDest.X = engine.bord.tileDest.Width * float32(i%engine.bord.mapW)
				engine.bord.tileDest.Y = engine.bord.tileDest.Height * float32(i/engine.bord.mapW)

				if engine.bord.srcMap[i] == "g" {
					engine.sprite.texture = engine.sprite.grass
				}
				if engine.bord.srcMap[i] == "i" {
					engine.sprite.texture = engine.sprite.hill
				}
				if engine.bord.srcMap[i] == "f" {
					engine.sprite.texture = engine.sprite.fence
					engine.bord.colisionList = append(engine.bord.colisionList, []float32{engine.bord.tileDest.X, engine.bord.tileDest.Y})
				}
				if engine.bord.srcMap[i] == "h" {
					engine.sprite.texture = engine.sprite.house
				}
				if engine.bord.srcMap[i] == "w" {
					engine.sprite.texture = engine.sprite.water
					engine.bord.colisionList = append(engine.bord.colisionList, []float32{engine.bord.tileDest.X, engine.bord.tileDest.Y})
				}
				if engine.bord.srcMap[i] == "t" {
					engine.sprite.texture = engine.sprite.tilled
				}

				if engine.bord.srcMap[i] == "h" || engine.bord.srcMap[i] == "f" { // si il y a une barrière ou une maison on met de l'herbe en dessous
					engine.bord.tileSrc.X = 0
					engine.bord.tileSrc.Y = 0
					rl.DrawTexturePro(engine.sprite.grass, engine.bord.tileSrc, engine.bord.tileDest, rl.NewVector2(engine.bord.tileDest.Width, engine.bord.tileDest.Height), 0, rl.White)
				}

				engine.bord.tileSrc.X = engine.bord.tileSrc.Width * float32((engine.bord.tileMap[i]-1)%int(engine.sprite.texture.Width/int32(engine.bord.tileSrc.Width)))
				engine.bord.tileSrc.Y = engine.bord.tileSrc.Height * float32((engine.bord.tileMap[i]-1)/int(engine.sprite.texture.Width/int32(engine.bord.tileSrc.Width)))
				rl.DrawTexturePro(engine.sprite.texture, engine.bord.tileSrc, engine.bord.tileDest, rl.NewVector2(engine.bord.tileDest.Width, engine.bord.tileDest.Height), 0, rl.White)
			}
		}

		rl.DrawTexturePro(engine.player.Sprite, engine.player.Src, engine.player.Dest, rl.NewVector2(engine.player.Dest.Width, engine.player.Dest.Height), 0, rl.White)
		rl.DrawText(strconv.Itoa(engine.character.hp), int32(engine.player.Dest.X)-218, int32(engine.player.Dest.Y)-134, 10, rl.Black)
		rl.DrawText(strconv.Itoa(engine.character.gold), int32(engine.player.Dest.X)+189, int32(engine.player.Dest.Y)-133, 10, rl.Black)
		rl.DrawTexture(engine.sprite.money, int32(engine.player.Dest.X)+199, int32(engine.player.Dest.Y)-138, rl.White)
		rl.DrawTexture(engine.sprite.heart, int32(engine.player.Dest.X)-238, int32(engine.player.Dest.Y)-138, rl.White)
		//rl.DrawTexture(engine.sprite.heartContainer, int32(engine.player.Dest.X) - 152, int32(engine.player.Dest.Y) - 78, rl.White)

		for i := 0; i < len(engine.monster); i++ {
			if engine.monster[i].alive {
				rl.DrawTexturePro(engine.monster[i].sprite, engine.monster[i].Src, engine.monster[i].Dest, rl.NewVector2(engine.monster[i].Dest.Width, engine.monster[i].Dest.Height), 0, rl.White)
			}
		}
	} else {
		rl.DrawTexture(engine.sprite.bgForest, int32(engine.player.Dest.X)-808, int32(engine.player.Dest.Y)-458, rl.White)
		rl.DrawText(strconv.Itoa(engine.character.hp), int32(engine.player.Dest.X)-800, int32(engine.player.Dest.Y)-430, 30, rl.Black)
		rl.DrawText(strconv.Itoa(engine.monster[engine.monsterBattle].hp), int32(engine.player.Dest.X)+600, int32(engine.player.Dest.Y)-400, 30, rl.Black)
		rl.DrawTexturePro(engine.monster[engine.monsterBattle].sprite, engine.monster[engine.monsterBattle].Src, engine.monster[engine.monsterBattle].Dest, rl.NewVector2(engine.monster[engine.monsterBattle].Dest.Width, engine.monster[engine.monsterBattle].Dest.Height), 0, rl.White)

		rl.DrawTexturePro(engine.sprite.buttonAttack, rl.NewRectangle(32, 32*6, 48, 32), rl.NewRectangle(engine.player.Dest.X-520, engine.player.Dest.Y+490, 230, 120), rl.NewVector2(200, 200), 0, rl.White)
		rl.DrawTexturePro(engine.sprite.buttonAttack, rl.NewRectangle(32+48, 32*6, 48, 32), rl.NewRectangle(engine.player.Dest.X-520, engine.player.Dest.Y+490, 230, 120), rl.NewVector2(200, 200), 0, rl.White)
		rl.DrawTexturePro(engine.sprite.buttonAttack, rl.NewRectangle(32*2+48*2, 32*3, 48, 32), rl.NewRectangle(engine.player.Dest.X+680, engine.player.Dest.Y+490, 230, 120), rl.NewVector2(200, 200), 0, rl.White)
		rl.DrawTexturePro(engine.sprite.buttonAttack, rl.NewRectangle(32*2+48*3, 32*3, 48, 32), rl.NewRectangle(engine.player.Dest.X+680, engine.player.Dest.Y+490, 230, 120), rl.NewVector2(200, 200), 0, rl.White)
		rl.DrawTexturePro(engine.sprite.invBar, rl.NewRectangle(0, 0, 203, 32), rl.NewRectangle(engine.player.Dest.X-180, engine.player.Dest.Y+495, 750, 112), rl.NewVector2(200, 200), 0, rl.White)
		for i := 0; i < len(engine.character.inventory); i++ {
			rl.DrawTexture(engine.character.inventory[i].sprite, (int32(engine.player.Dest.X) - engine.sprite.invBar.Width/2)+9+int32(19*i), (int32(engine.player.Dest.Y) - engine.sprite.invBar.Height/2)+8, rl.White)
		}
	}

	if engine.character.showInventory {
		rl.DrawTexture(engine.sprite.layer, int32(engine.player.Dest.X)-300, int32(engine.player.Dest.Y)-300, rl.White)
		rl.DrawTexture(engine.sprite.invBar, int32(engine.player.Dest.X)-engine.sprite.invBar.Width/2, int32(engine.player.Dest.Y)-engine.sprite.invBar.Height/2, rl.White)
		for i := 0; i < len(engine.character.inventory); i++ {
			rl.DrawTexture(engine.character.inventory[i].sprite, (int32(engine.player.Dest.X)-engine.sprite.invBar.Width/2)+9+int32(19*i), (int32(engine.player.Dest.Y)-engine.sprite.invBar.Height/2)+8, rl.White)
		}
	}
}
