package main

import "github.com/gen2brain/raylib-go/raylib"

func update() {
	run = !rl.WindowShouldClose()

	playerSrc.X = playerSrc.Width * float32(playerFrame)

	if playerMoving {
		if playerUp { playerDest.Y -= playerSpeed }
		if playerDown { playerDest.Y += playerSpeed }
		if playerLeft { playerDest.X -= playerSpeed }
		if playerRight { playerDest.X += playerSpeed }
		if framCount % 8 == 1 { playerFrame++ }
	} else if framCount % 45 == 1 {
		playerFrame++
	}

	framCount++
	if playerFrame > 3 { playerFrame = 0 }
	if !playerMoving && playerFrame > 1 { playerFrame = 0 }

	playerSrc.X = playerSrc.Width * float32(playerFrame)
	playerSrc.Y = playerSrc.Height * float32(playerDir)

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(float32(playerDest.X - (playerDest.Width / 2)), float32(playerDest.Y - (playerDest.Height/2)))

	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false
}
