package main

import "github.com/gen2brain/raylib-go/raylib"

func update() {
	run = !rl.WindowShouldClose()

	player.playerSrc.X = player.playerSrc.Width * float32(player.playerFrame)

	if player.playerMoving {
		if player.playerUp { player.playerDest.Y -= player.playerSpeed }
		if player.playerDown { player.playerDest.Y += player.playerSpeed }
		if player.playerLeft { player.playerDest.X -= player.playerSpeed }
		if player.playerRight { player.playerDest.X += player.playerSpeed }
		if framCount % 8 == 1 { player.playerFrame++ }
	} else if framCount % 45 == 1 {
		player.playerFrame++
	}

	framCount++
	if player.playerFrame > 3 { player.playerFrame = 0 }
	if !player.playerMoving && player.playerFrame > 1 { player.playerFrame = 0 }

	player.playerSrc.X = player.playerSrc.Width * float32(player.playerFrame)
	player.playerSrc.Y = player.playerSrc.Height * float32(player.playerDir)

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(float32(player.playerDest.X - (player.playerDest.Width / 2)), float32(player.playerDest.Y - (player.playerDest.Height/2)))

	player.playerMoving = false
	player.playerUp, player.playerDown, player.playerRight, player.playerLeft = false, false, false, false
}
