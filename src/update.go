package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func update() {
	run = !rl.WindowShouldClose()

	player.Src.Y = player.Src.Width * float32(player.Frame)

	for i := range bord.colisionList {
		if player.Up {
			nextBlock := []float32{player.Dest.X, player.Dest.Y+player.Src.Height-17}
			if (bord.colisionList[i][0] <= nextBlock[0] && nextBlock[0] <= bord.colisionList[i][0]+16) && (bord.colisionList[i][1] <= nextBlock[1] && nextBlock[1] <= bord.colisionList[i][1]+16) {
				player.Moving = false
			}
		}
		if player.Down {
			nextBlock := []float32{player.Dest.X, player.Dest.Y+18}
			if (bord.colisionList[i][0] <= nextBlock[0] && nextBlock[0] <= bord.colisionList[i][0]+16) && (bord.colisionList[i][1] <= nextBlock[1] && nextBlock[1] <= bord.colisionList[i][1]+16) {
				player.Moving = false
			}
		}
		if player.Left {
			nextBlock := []float32{player.Dest.X+player.Src.Width-17, player.Dest.Y}
			if (bord.colisionList[i][0] <= nextBlock[0] && nextBlock[0] <= bord.colisionList[i][0]+16) && (bord.colisionList[i][1] <= nextBlock[1] && nextBlock[1] <= bord.colisionList[i][1]+16) {
				player.Moving = false
			}
		}
		if player.Right {
			nextBlock := []float32{player.Dest.X+16, player.Dest.Y}
			if (bord.colisionList[i][0] <= nextBlock[0] && nextBlock[0] <= bord.colisionList[i][0]+16) && (bord.colisionList[i][1] <= nextBlock[1] && nextBlock[1] <= bord.colisionList[i][1]+16) {
				player.Moving = false
			}
		}
	}

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
