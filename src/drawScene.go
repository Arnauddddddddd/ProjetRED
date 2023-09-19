package main

import (
	"github.com/gen2brain/raylib-go/raylib"

	"strconv"
)

func drawScene() {
	for i := 0; i < len(bord.tileMap); i++ {
		if bord.tileMap[i] != 0 {
			bord.tileDest.X = bord.tileDest.Width * float32(i % bord.mapW)
			bord.tileDest.Y = bord.tileDest.Height * float32(i / bord.mapW)
			
			if bord.srcMap[i] == "g" { texture = sprite.grass }
			if bord.srcMap[i] == "i" { texture = sprite.hill }
			if bord.srcMap[i] == "f" { texture = sprite.fence }
			if bord.srcMap[i] == "h" { texture = sprite.house }
			if bord.srcMap[i] == "w" { 
				texture = sprite.water
				//stop = append(stop, []float32{bord.tileDest.X, bord.tileDest.Y})
			}
			if bord.srcMap[i] == "t" { texture = sprite.tilled }

			if bord.srcMap[i] == "h" || bord.srcMap[i] == "f" { // si il y a une barrière ou une maison on met de l'herbe en dessous
				bord.tileSrc.X = 0
				bord.tileSrc.Y = 0
				rl.DrawTexturePro(sprite.grass, bord.tileSrc, bord.tileDest, rl.NewVector2(bord.tileDest.Width, bord.tileDest.Height), 0, rl.White)
			}
			
			bord.tileSrc.X = bord.tileSrc.Width * float32((bord.tileMap[i]-1) % int(texture.Width / int32(bord.tileSrc.Width)))
			bord.tileSrc.Y = bord.tileSrc.Height * float32((bord.tileMap[i]-1) / int(texture.Width / int32(bord.tileSrc.Width)))
			rl.DrawTexturePro(texture, bord.tileSrc, bord.tileDest, rl.NewVector2(bord.tileDest.Width, bord.tileDest.Height), 0, rl.White)
		}
	}
	rl.DrawTexturePro(player.Sprite, player.Src, player.Dest, rl.NewVector2(player.Dest.Width, player.Dest.Height), 0, rl.White)
    rl.DrawText(strconv.Itoa(int(rl.GetMouseX())) + " " + strconv.Itoa(int(rl.GetMouseY())), 0, 0, 20, rl.Black)	
}
