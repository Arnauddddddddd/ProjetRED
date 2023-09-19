package main

import (
	//"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

func update() {
	run = !rl.WindowShouldClose()

	player.Src.Y = player.Src.Width * float32(player.Frame)

	if player.Moving {
		if player.Up { player.Dest.Y -= player.Speed }
		if player.Down { player.Dest.Y += player.Speed }
		if player.Left { player.Dest.X -= player.Speed }
		if player.Right { player.Dest.X += player.Speed }
		if framCount % 8 == 1 { player.Frame++ }
	}

	framCount++
	if player.Frame > 3 { player.Frame = 0 }
	if !player.Moving && player.Frame > 1 { player.Frame = 0 }

	player.Src.Y = player.Src.Width * float32(player.Frame)
	player.Src.X = player.Src.Height * float32(player.Dir)

	//fmt.Println(player.Dest.X)

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(float32(player.Dest.X - (player.Dest.Width / 2)), float32(player.Dest.Y - (player.Dest.Height/2)))

	player.Moving = false
	player.Up, player.Down, player.Right, player.Left = false, false, false, false
}
