package main

import (
	"github.com/gen2brain/raylib-go/raylib"

	"strconv"
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(player.playerSprite, player.playerSrc, player.playerDest, rl.NewVector2(player.playerDest.Width, player.playerDest.Height), 0, rl.White)
    rl.DrawText(strconv.Itoa(int(rl.GetMouseX())) + " " + strconv.Itoa(int(rl.GetMouseY())), 0, 0, 20, rl.Black)	
}
