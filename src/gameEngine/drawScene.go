package gameEngine

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawScene(engine *EngineStruct) {
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
			rl.DrawTexturePro(engine.monster[i].sprite, rl.NewRectangle(engine.monster[i].Src.X+float32(engine.monster[i].Xstart), engine.monster[i].Src.Y, engine.monster[i].Src.Width, engine.monster[i].Src.Height), engine.monster[i].Dest, rl.NewVector2(engine.monster[i].Dest.Width, engine.monster[i].Dest.Height), 0, rl.White)
		}
	}

	rl.DrawTexturePro(engine.shopKeeper.sprite, engine.shopKeeper.Src, engine.shopKeeper.Dest, rl.NewVector2(engine.shopKeeper.Dest.Width, engine.shopKeeper.Dest.Height), 0, rl.White)

	if engine.character.showInventory {
		rl.DrawTexture(engine.sprite.layer, int32(engine.player.Dest.X)-300, int32(engine.player.Dest.Y)-300, rl.White)
		rl.DrawTexture(engine.sprite.invBar, int32(engine.player.Dest.X)-engine.sprite.invBar.Width/2 - int32(engine.player.Dest.Width / 2), int32(engine.player.Dest.Y)-engine.sprite.invBar.Height/2- int32(engine.player.Dest.Height / 2), rl.White)
		t := 0
		slot2 := 0
		for i := 0; i < len(engine.character.inventory); i++ {
			if i == 4 {
				slot2 = 1
				t = 0
			}
			rl.DrawTexture(engine.character.inventory[i].sprite, (int32(engine.player.Dest.X)-engine.sprite.invBar.Width/2)+10+int32(23*t)- int32(engine.player.Dest.Width / 2), (int32(engine.player.Dest.Y)-engine.sprite.invBar.Height/2)+10+int32(23*slot2)- int32(engine.player.Dest.Height / 2), rl.White)
			t++
		}
	}

	if engine.shop && engine.character.showInventory {
		rl.DrawTexture(engine.sprite.layer, int32(engine.player.Dest.X)-300, int32(engine.player.Dest.Y)-300, rl.White)
		rl.DrawTexture(engine.shopKeeper.shopSprite, int32(engine.player.Dest.X)-engine.sprite.invBar.Width/2 - int32(engine.player.Dest.Width / 2) + 12, int32(engine.player.Dest.Y)-engine.sprite.invBar.Height/2- int32(engine.player.Dest.Height / 2) - 100, rl.White)
		t := 0
		slot2 := 0
		for i := 0; i < len(engine.shopKeeper.items); i++ {
			if i == 3 || i == 6 {
				slot2 += 1
				t = 0
			}
			rl.DrawTexture(engine.shopKeeper.items[i].sprite, (int32(engine.player.Dest.X)-engine.sprite.invBar.Width/2)+10+int32(23*t) - int32(engine.player.Dest.Width / 2) + 12, (int32(engine.player.Dest.Y)-engine.sprite.invBar.Height/2)+10+int32(23*slot2)- int32(engine.player.Dest.Height / 2) - 100, rl.White)
			t++
		}
		rl.DrawTexturePro(engine.sprite.shopName, rl.NewRectangle(0, 0, 213, 101), rl.NewRectangle(float32(engine.player.Dest.X+25), float32(engine.player.Dest.Y+33)-53, 100, 50), rl.NewVector2(float32(engine.sprite.shopName.Width), float32(engine.sprite.shopName.Height)), 0, rl.White)
		rl.DrawText(strconv.Itoa(engine.character.gold), int32(engine.player.Dest.X)+189, int32(engine.player.Dest.Y)-133, 10, rl.Black)
		rl.DrawTexture(engine.sprite.money, int32(engine.player.Dest.X)+199, int32(engine.player.Dest.Y)-138, rl.White)
	}
}
