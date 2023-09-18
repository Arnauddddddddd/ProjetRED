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
			
			if bord.srcMap[i] == "g" { texture = grassSprite }
			
			bord.tileSrc.X = bord.tileSrc.Width * float32((bord.tileMap[i]-1) % int(texture.Width / int32(bord.tileSrc.Width)))
			bord.tileSrc.Y = bord.tileSrc.Height * float32((bord.tileMap[i]-1) / int(texture.Width / int32(bord.tileSrc.Width)))
			rl.DrawTexturePro(texture, bord.tileSrc, bord.tileDest, rl.NewVector2(bord.tileDest.Width, bord.tileDest.Height), 0, rl.White)
		}
	}
	rl.DrawTexturePro(player.playerSprite, player.playerSrc, player.playerDest, rl.NewVector2(player.playerDest.Width, player.playerDest.Height), 0, rl.White)
    rl.DrawText(strconv.Itoa(int(rl.GetMouseX())) + " " + strconv.Itoa(int(rl.GetMouseY())), 0, 0, 20, rl.Black)	
}
