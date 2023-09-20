package main

import (
	"strconv"

	"github.com/gen2brain/raylib-go/raylib"
)

func drawScene() {
	bord.colisionList = [][]float32{}
	for i := 0; i < len(bord.tileMap); i++ {
		if bord.tileMap[i] != 0 {
			bord.tileDest.X = bord.tileDest.Width * float32(i % bord.mapW)
			bord.tileDest.Y = bord.tileDest.Height * float32(i / bord.mapW)
			
			if bord.srcMap[i] == "g" { sprite.texture = sprite.grass }
			if bord.srcMap[i] == "i" { sprite.texture = sprite.hill }
			if bord.srcMap[i] == "f" {
				sprite.texture = sprite.fence
				bord.colisionList = append(bord.colisionList, []float32{bord.tileDest.X, bord.tileDest.Y})
			}
			if bord.srcMap[i] == "h" { sprite.texture = sprite.house }
			if bord.srcMap[i] == "w" { 
				sprite.texture = sprite.water
				bord.colisionList = append(bord.colisionList, []float32{bord.tileDest.X, bord.tileDest.Y})
			}
			if bord.srcMap[i] == "t" { sprite.texture = sprite.tilled }

			if bord.srcMap[i] == "h" || bord.srcMap[i] == "f" { // si il y a une barrière ou une maison on met de l'herbe en dessous
				bord.tileSrc.X = 0
				bord.tileSrc.Y = 0
				rl.DrawTexturePro(sprite.grass, bord.tileSrc, bord.tileDest, rl.NewVector2(bord.tileDest.Width, bord.tileDest.Height), 0, rl.White)
			}
			
			bord.tileSrc.X = bord.tileSrc.Width * float32((bord.tileMap[i]-1) % int(sprite.texture.Width / int32(bord.tileSrc.Width)))
			bord.tileSrc.Y = bord.tileSrc.Height * float32((bord.tileMap[i]-1) / int(sprite.texture.Width / int32(bord.tileSrc.Width)))
			rl.DrawTexturePro(sprite.texture, bord.tileSrc, bord.tileDest, rl.NewVector2(bord.tileDest.Width, bord.tileDest.Height), 0, rl.White)
		}
	}

	rl.DrawTexturePro(player.Sprite, player.Src, player.Dest, rl.NewVector2(player.Dest.Width, player.Dest.Height), 0, rl.White)
	rl.DrawText(strconv.Itoa(character.hp), int32(player.Dest.X) - 218, int32(player.Dest.Y) - 134, 10, rl.Black)
	rl.DrawText(strconv.Itoa(character.gold), int32(player.Dest.X) + 189, int32(player.Dest.Y) - 133, 10, rl.Black)
	rl.DrawTexture(sprite.money, int32(player.Dest.X) + 199, int32(player.Dest.Y) - 138, rl.White)
	rl.DrawTexture(sprite.heart, int32(player.Dest.X) - 238, int32(player.Dest.Y) - 138, rl.White)
	//rl.DrawTexture(sprite.heartContainer, int32(player.Dest.X) - 152, int32(player.Dest.Y) - 78, rl.White)

	for i := 0; i < len(monster); i++ {
		rl.DrawTexturePro(monster[i].sprite, monster[i].Src, monster[i].Dest, rl.NewVector2(monster[i].Dest.Width, monster[i].Dest.Height), 0, rl.White)
	}

	if character.showInventory {
		rl.DrawTexture(sprite.layer, int32(player.Dest.X) - 300, int32(player.Dest.Y) - 300, rl.White)
		rl.DrawTexture(sprite.invBar, int32(player.Dest.X) - sprite.invBar.Width/2, int32(player.Dest.Y) - sprite.invBar.Height/2, rl.White)
		for i := 0; i < len(character.inventory); i++ {
			rl.DrawTexture(character.inventory[i].sprite, (int32(player.Dest.X) - sprite.invBar.Width/2)+9+int32(19*i), (int32(player.Dest.Y) - sprite.invBar.Height/2)+8, rl.White)
		}
	}
}
